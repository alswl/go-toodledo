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

	tasksPane            taskspane.Model
	sidebar              sidebar.Model
	statusBar            statusbar.Model
	focusableModelsIndex []string

	//activateModel string
	focusedIndex int

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
		if m.isInputting {
			// XXX
			//m.Focused().HandleKey(msg)
			return m, nil
		} else {
			switch msg.String() {
			case "tab":
				// TODO refactor, switch with keymap
				// loopFocus change the model fields(isFocused)
				m.loopFocus()
				return m, nil
			}
		}

		// 3. sub component
		newM, cmd := m.getFocusedModel().Update(msg)
		switch newM.(type) {
		case taskspane.Model:
			m.tasksPane = newM.(taskspane.Model)
			return m, cmd
		case sidebar.Model:
			m.sidebar = newM.(sidebar.Model)
			return m, cmd
		case statusbar.Model:
			m.statusBar = newM.(statusbar.Model)
			return m, cmd
		}

		// normal keypress
	case tea.WindowSizeMsg:
		sideBarWidth := msg.Width / 12 * 3
		m.sidebar.Resize(sideBarWidth, msg.Height-1)
		m.tasksPane.Resize(msg.Width-sideBarWidth, msg.Height-1)
		m.statusBar.SetSize(msg.Width)
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

func (m *Model) loopFocus() {
	m.getFocusedModelF().Blur()
	m.focusedIndex = (m.focusedIndex + 1) % len(m.focusableModelsIndex)
	m.getFocusedModelF().Focus()
}

func (m *Model) getFocusedModelF() components.FocusableInterface {
	name := m.focusableModelsIndex[m.focusedIndex]
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
	name := m.focusableModelsIndex[m.focusedIndex]
	switch name {
	case "tasks":
		return m.tasksPane
	case "sidebar":
		return m.sidebar
	case "statusbar":
		return m.statusBar

	}
	panic("unreachable")
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
	// XXX
	statusB.SetContent("tasks", "a fox jumped over the lazy dog", "1/999", "PAUSE")

	tp := taskspane.InitModel(tasks)
	sb := sidebar.InitModel()
	m := Model{
		tasksPane: tp,
		sidebar:   sb,
		statusBar: statusB,
		focusableModelsIndex: []string{
			"tasks",
			"sidebar",
			//"statusB",
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
	m.focusedIndex = 0
	m.tasksPane.Focus()

	return m
}
