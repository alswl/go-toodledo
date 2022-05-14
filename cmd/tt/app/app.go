package app

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/cmd/tt/components"
	"github.com/alswl/go-toodledo/cmd/tt/components/sidebar"
	"github.com/alswl/go-toodledo/cmd/tt/components/statusbar"
	"github.com/alswl/go-toodledo/cmd/tt/components/taskspane"
	"github.com/alswl/go-toodledo/cmd/tt/mockdata"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/cmd/tt/utils"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"os"
)

type States struct {
	// Tasks is available tasks
	Tasks []*models.RichTask
	// XXX
	Filter string
}

type Model struct {
	keys utils.KeyMap
	err  error
	//config
	data   []*models.RichTask
	states *States

	tasksPane taskspane.Model
	sidebar   sidebar.Model
	statusBar statusbar.Model

	allModels []string
	focused   string

	tabSupportModels []string
	tabIndex         int

	// TODO help pane
	//help          help.Model
	// TODO ready check
	ready         bool
	isSidebarOpen bool
	width         int
	//filterWindow  FilterFormModel
	//tabs            tabs.Model
	//context         context.ProgramContext
	isInputting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// process logics
	// 1. global keymap
	// 2. input TODO input is one of focuesd component?
	// 3. focused component

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// TODO keymap
		// 1. global keymap
		if funk.ContainsString([]string{"ctrl+c"}, msg.String()) {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			}
		}

		// 2. app mode
		if !m.isInputting {
			switch msg.String() {
			case "tab":
				// TODO refactor, switch with keymap
				// loopTabFocus change the model fields(isFocused)
				m.loopTabFocus()
				return m, nil
			}
		}

		// 3. sub component
		return m.updateFocusedModel(msg)

		// normal keypress
	case tea.WindowSizeMsg:
		sideBarWidth := msg.Width / 12 * 3
		m.sidebar.Resize(sideBarWidth, msg.Height-1)
		m.tasksPane.Resize(msg.Width-sideBarWidth, msg.Height-1)
		m.statusBar.Resize(msg.Width, 0)
	}
	return m, nil
}

func (m Model) View() string {
	// TODO error handling
	if m.err != nil {
		return m.err.Error()
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

func (m Model) GetStates() States {
	return *m.states
}

func (m *Model) loopTabFocus() {
	if !funk.ContainsString(m.tabSupportModels, m.focused) {
		return
	}

	//m.getFocusedModelF().Blur()
	//old := m.tabSupportModels[m.tabIndex]
	m.tabIndex = (m.tabIndex + 1) % len(m.tabSupportModels)
	new := m.tabSupportModels[m.tabIndex]
	m.focus(new)
}

func (m *Model) getFocusedModelF() components.FocusableInterface {
	name := m.tabSupportModels[m.tabIndex]
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

func (m Model) updateFocusedModel(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
	case sidebar.Model:
		newM, cmd = mm.Update(msg)
		m.sidebar = newM.(sidebar.Model)
		return m, cmd
	case statusbar.Model:
		newM, cmd = mm.Update(msg)
		m.statusBar = newM.(statusbar.Model)
		m.tasksPane.Filter(m.statusBar.GetFilterInput())

		// post action
		switch msg.String() {
		case "enter":
			m.isInputting = false
			m.focus("tasks")
		}
		return m, cmd
	}
	return m, cmd
}

//func (m *Model) SetInputting(is bool) {
//	m.isInputting = is
//}

//func (m Model) IsInputting() bool {
//	return m.isInputting
//}

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

func InitialModel() Model {
	var err error
	var tasks []*models.RichTask
	if os.Getenv("MOCK") == "true" {
		tasks, err = mockdata.AllTasksMock()
	} else {
		tasks, err = AllTasks()
	}
	if err != nil {
		// FIXME
		panic(err)
	}

	// FIXME move tables km to model
	keys := table.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down")
	keys.RowUp.SetKeys("k", "up")

	statusB := statusbar.NewDefault()
	// XXX ?
	statusB.SetContent("tasks", "a fox jumped over the lazy dog", "1/999", "PAUSE")

	// FIXME tasks should comes from syncer
	tp := taskspane.InitModel(tasks)
	sb := sidebar.InitModel()
	m := Model{
		tasksPane: tp,
		sidebar:   sb,
		statusBar: statusB,
		tabSupportModels: []string{
			"tasks",
			"sidebar",
			//"statusB",
		},
		allModels: []string{
			"tasks",
			"sidebar",
			"statusbar",
		},
		//allFocusableModels: []components.FocusableInterface{
		//	&tp,
		//	&sb,
		//},
		//allModels: []tea.Model{
		//	&tp,
		//	&sb,
		//},
	}
	// FIXME focus as an method
	//m.av = "tasks"
	m.tabIndex = 0
	m.focused = "tasks"
	m.tasksPane.Focus()

	return m
}
