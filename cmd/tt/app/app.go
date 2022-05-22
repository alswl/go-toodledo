package app

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/cmd/tt/components"
	comsidebar "github.com/alswl/go-toodledo/cmd/tt/components/sidebar"
	comstatusbar "github.com/alswl/go-toodledo/cmd/tt/components/statusbar"
	"github.com/alswl/go-toodledo/cmd/tt/components/taskspane"
	"github.com/alswl/go-toodledo/cmd/tt/mockdata"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

var allModels = []string{
	"tasks",
	"sidebar",
	"statusbar",
}
var supportModels = []string{
	"tasks",
	"sidebar",
	//"statusB",
}

type States struct {
	// Tasks is available tasks
	Tasks    []*models.RichTask
	Contexts []models.Context
	// XXX
	Filter string
}

type Model struct {
	// properties
	// TODO

	// states TODO
	states   *States
	err      error
	focused  string
	tabIndex int
	// TODO ready check
	ready bool
	//isSidebarOpen bool

	// view
	tasksPane taskspane.Model
	sidebar   comsidebar.Model
	statusBar comstatusbar.Model
	// TODO help pane
	//help          help.Model
	isInputting bool
}

func (m *Model) Init() tea.Cmd {
	contextSvc, err := injector.InitContextCachedService()
	if err != nil {
		m.err = err
		return nil
	}
	cs, err := contextSvc.ListAll()
	if err != nil {
		m.err = err
		return nil
	}
	m.states.Contexts = utils.UnwrapListPointer(cs)
	m.sidebar, _ = m.sidebar.UpdateX(utils.UnwrapListPointer(cs))

	m.initTasks()
	m.tasksPane, _ = m.tasksPane.UpdateX(m.states.Tasks)

	return nil
}

// Update the Model states, this Model is only one instance and global using
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// process logics
	// 1. global keymap
	// 2. input TODO input is one of focused component?
	// 3. focused component

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		sideBarWidth := msg.Width / 12 * 3
		m.sidebar.Resize(sideBarWidth, msg.Height-1)
		m.tasksPane.Resize(msg.Width-sideBarWidth, msg.Height-1)
		m.statusBar.Resize(msg.Width, 0)
	case tea.KeyMsg: // handle event bubble
		// 1. global keymap
		if funk.ContainsString([]string{"ctrl+c"}, msg.String()) {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			}
		}

		// 2. main app
		if !m.isInputting {
			switch msg.String() {
			case "tab":
				// change the model fields(isFocused)
				m.loopFocus()
				return m, cmd
			}
		}

		// 3. sub focused component
		m, cmd = m.updateFocusedModel(msg)

		// normal keypress
	}
	return m, cmd
}

func (m *Model) View() string {
	if m.err != nil {
		m.statusBar.SetMode("ERROR")
		m.statusBar.SetStatus(m.err.Error())
	}

	return styles.MainPaneStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				m.sidebar.View(),
				m.tasksPane.View(),
			),
			m.statusBar.View(),
		),
	)
}

func (m *Model) loopFocus() {
	if !funk.ContainsString(supportModels, m.focused) {
		return
	}

	m.tabIndex = (m.tabIndex + 1 + len(supportModels)) % len(supportModels)
	new := supportModels[m.tabIndex]
	m.focus(new)
}

func (m *Model) getFocusedModelF() components.FocusableInterface {
	name := supportModels[m.tabIndex]
	switch name {
	case "tasks":
		return &m.tasksPane
	case "sidebar":
		return &m.sidebar
	case "statusbar":
		return &m.statusBar

	}
	panic("unreachable")
}

func (m *Model) getFocusedModel() tea.Model {
	switch m.focused {
	case "tasks":
		return m.tasksPane
	case "sidebar":
		return m.sidebar
	case "statusbar":
		return m.statusBar

	}
	panic("unreachable")
}

func (m *Model) focus(name string) {
	m.getFocusedModelF().Blur()
	m.focused = name
	m.getFocusedModelF().Focus()
}

func (m *Model) focusStatusBar() {
	m.focus("statusbar")
}

func (m *Model) updateFocusedModel(msg tea.KeyMsg) (*Model, tea.Cmd) {
	var newM tea.Model
	var cmd tea.Cmd
	mm := m.getFocusedModel()

	// post action in main app
	switch mm.(type) {
	case taskspane.Model:
		switch msg.String() {
		case "/":
			m.isInputting = true
			m.focusStatusBar()
			m.statusBar.FocusFilter()
		default:
			newM, cmd = mm.Update(msg)
			m.tasksPane = newM.(taskspane.Model)
		}
		return m, cmd
	case comsidebar.Model:
		newM, cmd = mm.Update(msg)
		m.sidebar = newM.(comsidebar.Model)
		return m, cmd
	case comstatusbar.Model:
		newM, cmd = mm.Update(msg)
		m.statusBar = newM.(comstatusbar.Model)
		m.tasksPane.Filter(m.statusBar.GetFilterInput())

		// post action
		switch msg.String() {
		case "enter", "esc":
			m.isInputting = false
			m.focus("tasks")
		}
		return m, cmd
	}
	return m, cmd
}

func (m *Model) OnItemChange(tab string, item comsidebar.Item) error {
	m.statusBar.SetStatus("tab: " + tab + " item: " + item.Title())
	svc, err := injector.InitTaskCachedService()
	if err != nil {
		m.statusBar.SetStatus("ERROR: " + err.Error())
	}
	taskRichSvc, err := injector.InitTaskRichService()
	if err != nil {
		m.statusBar.SetStatus("ERROR: " + err.Error())
	}
	query := &queries.TaskListQuery{}
	if tab == "Contexts" {
		query.ContextID = item.ID()
	}
	tasks, err := svc.ListAllByQuery(query)
	if err != nil {
		m.statusBar.SetStatus("ERROR: " + err.Error())
	}
	rts, _ := taskRichSvc.RichThem(tasks)
	m.states.Tasks = rts
	m.tasksPane, _ = m.tasksPane.UpdateX(m.states.Tasks)
	m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

	return nil
}

func (m Model) initTasks() {
	var tasks []*models.RichTask
	var err error
	//if os.Getenv("MOCK") == "true" {
	tasks, err = mockdata.AllTasksMock()
	//} else {
	//	// FIXME query tasks with filter
	//	tasks, err = AllTasks()
	//}
	if err != nil {
		m.err = err
		return
	}

	m.states.Tasks = tasks
}

// FIXME using daemon syncer
// nolint:deadcode
func AllTasks() ([]*models.RichTask, error) {
	_, err := injector.InitApp()
	if err != nil {
		logrus.Fatal("login required, using `toodledo auth login` to login.")
		return nil, err
	}
	svc, err := injector.InitTaskCachedService()
	if err != nil {
		logrus.WithError(err).Fatal("failed to init task service")
		return nil, err
	}
	syncer, err := injector.InitSyncer()
	if err != nil {
		logrus.WithError(err).Fatal("init syncer failed")
		return nil, err
	}
	taskRichSvc, err := injector.InitTaskRichService()
	if err != nil {
		logrus.WithError(err).Fatal("init task rich service failed")
		return nil, err
	}
	err = syncer.SyncOnce()
	if err != nil {
		logrus.WithError(err).Fatal("sync failed")
		return nil, err
	}
	tasks, err := svc.ListAllByQuery(&queries.TaskListQuery{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	rts, err := taskRichSvc.RichThem(tasks)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rts, nil
}

func InitialModel() *Model {
	var err error
	if err != nil {
		// FIXME
		panic(err)
	}
	_, err = injector.InitApp()
	if err != nil {
		panic(err)
	}

	states := &States{}

	statusBar := comstatusbar.NewDefault()
	statusBar.SetMode("tasks")
	statusBar.SetStatus("a fox jumped over the lazy dog")
	statusBar.SetInfo1("1/999")
	statusBar.SetInfo2("HELP")

	// FIXME tasks should comes from syncer
	taskPane := taskspane.InitModel(states.Tasks)

	m := Model{
		tasksPane: taskPane,
		statusBar: statusBar,
		states:    states,
	}
	sb := comsidebar.InitModel(comsidebar.Properties{}, m.OnItemChange)
	m.sidebar = sb
	// FIXME focus as an method
	m.tabIndex = 0
	m.focused = "tasks"
	m.tasksPane.Focus()

	return &m
}
