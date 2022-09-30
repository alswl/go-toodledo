package app

import (
	"context"
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/cmd/tt/components"
	comsidebar "github.com/alswl/go-toodledo/cmd/tt/components/sidebar"
	comstatusbar "github.com/alswl/go-toodledo/cmd/tt/components/statusbar"
	"github.com/alswl/go-toodledo/cmd/tt/components/taskspane"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/fetchers"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"time"
)

var (
	allModels = []string{
		"tasks",
		"sidebar",
		"statusbar",
	}
	supportModels = []string{
		"tasks",
		"sidebar",
	}
)

type States struct {
	// Tasks is available tasks
	Tasks    []*models.RichTask
	Contexts []*models.Context
	Folders  []*models.Folder
	Goals    []*models.Goal
	query    *queries.TaskListQuery
}

// Model is the main tt app
type Model struct {
	taskRichSvc  services.TaskRichService
	contextSvc   services.ContextService
	folderSvc    services.FolderService
	goalSvc      services.GoalService
	taskSvc      services.TaskExtendedService
	taskLocalSvc services.TaskLocalService

	// properties
	log logrus.FieldLogger

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
	fetcher     fetchers.DaemonFetcher
}

func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, func() tea.Msg {
		cs, err := m.contextSvc.ListAll()
		if err != nil {
			m.err = err
			return nil
		}
		fs, err := m.folderSvc.ListAll()
		if err != nil {
			m.err = err
			return nil
		}
		gs, err := m.goalSvc.ListAll()
		if err != nil {
			m.err = err
		}

		// Contexts are first tab in sidebar
		m.states.Contexts = cs
		m.states.Contexts = append([]*models.Context{{
			ID:   0,
			Name: "All",
		}}, cs...)
		m.states.Contexts = append(m.states.Contexts, &models.Context{
			ID:   -1,
			Name: "None",
		})
		m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Contexts))
		// using default first now
		m.states.query.ContextID = m.states.Contexts[0].ID

		// folders
		m.states.Folders = fs
		m.states.Folders = append([]*models.Folder{{
			ID:   0,
			Name: "All",
		}}, fs...)
		m.states.Folders = append(m.states.Folders, &models.Folder{
			ID:   -1,
			Name: "None",
		})
		m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Folders))

		// goals
		m.states.Goals = gs
		m.states.Goals = append([]*models.Goal{{
			ID:   0,
			Name: "All",
		}}, gs...)
		m.states.Goals = append(m.states.Goals, &models.Goal{
			ID:   -1,
			Name: "None",
		})
		m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Goals))

		// tasks
		tasks, err := m.taskSvc.ListAllByQuery(m.states.query)
		if err != nil {
			m.statusBar.SetStatus("ERROR: " + err.Error())
		}
		rts, _ := m.taskRichSvc.RichThem(tasks)
		m.states.Tasks = rts
		m.tasksPane, _ = m.tasksPane.UpdateTyped(m.states.Tasks)
		m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

		return nil
	})
	cmds = append(cmds, func() tea.Msg {
		m.fetcher.Start(context.Background())
		return nil
	})

	return tea.Batch(cmds...)
}

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

		// 2. main app, command mode
		if !m.isInputting {
			switch msg.String() {
			case "tab":
				// change the model fields(isFocused)
				m.loopFocus()
				return m, cmd
			case "r":
				// FIXME ui refresh, message logic fix
				err := m.fetcher.Notify()
				if err != nil {
					m.log.WithError(err).Error("notify fetcher")
				}
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
		mmTyped := mm.(comsidebar.Model)
		m.sidebar, cmd = mmTyped.UpdateTyped(msg)
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
	m.states.query = &queries.TaskListQuery{}
	switch tab {
	case constants.Contexts:
		m.states.query.ContextID = item.ID()
	case constants.Folders:
		m.states.query.FolderID = item.ID()
	case constants.Goals:
		m.states.query.GoalID = item.ID()
	}
	// FIXME using taskSvc
	tasks, err := m.taskLocalSvc.ListAllByQuery(m.states.query)
	if err != nil {
		m.statusBar.SetStatus("ERROR: " + err.Error())
	}
	rts, _ := m.taskRichSvc.RichThem(tasks)
	m.states.Tasks = rts
	m.tasksPane, _ = m.tasksPane.UpdateTyped(m.states.Tasks)
	m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

	return nil
}

func InitialModel() *Model {
	// prepare
	log := logging.GetLogger("tt")
	var err error
	app, err := injector.InitTUIApp()
	if err != nil {
		// TODO
		panic(err)
	}
	taskSvc := app.TaskLocalSvc
	taskLocalSvc := app.TaskLocalSvc
	taskRichSvc := app.TaskRichSvc
	contextSvc := app.ContextLocalSvc
	folderSvc := app.FolderLocalSvc
	goalSvc := app.GoalLocalSvc
	states := &States{
		Tasks:    []*models.RichTask{},
		Contexts: []*models.Context{},
		Folders:  []*models.Folder{},
		Goals:    []*models.Goal{},
		query:    &queries.TaskListQuery{},
	}

	// status bar
	statusBar := comstatusbar.NewDefault()
	statusBar.SetMode("tasks")
	statusBar.SetStatus("a fox jumped over the lazy dog")
	statusBar.SetInfo1("1/999")
	statusBar.SetInfo2("HELP")

	// task pane
	taskPane := taskspane.InitModel(states.Tasks)

	sidebar := comsidebar.InitModel(comsidebar.Properties{})

	// main app
	m := Model{
		log:          log,
		taskRichSvc:  taskRichSvc,
		contextSvc:   contextSvc,
		folderSvc:    folderSvc,
		goalSvc:      goalSvc,
		taskSvc:      taskSvc,
		taskLocalSvc: taskLocalSvc,
		states:       states,
		err:          nil,
		focused:      "tasks",
		tabIndex:     0,
		ready:        false,
		tasksPane:    taskPane,
		statusBar:    statusBar,
		sidebar:      sidebar,
		isInputting:  false,
	}

	m.sidebar.RegisterItemChange(m.OnItemChange)

	// init fetcher
	fetcher := fetchers.NewSimpleFetcher(log, 1*time.Minute, fetchers.NewToodledoFetchFnPartial(
		log,
		app.FolderLocalSvc,
		app.ContextLocalSvc,
		app.GoalLocalSvc,
		app.TaskLocalSvc,
		app.AccountSvc,
	), fetchers.NewStatusDescriber(func() error {
		// TODO using register fun instead of invoke m in New func
		m.statusBar.SetStatus("fetching...")
		return nil
	}, func() error {
		// TODO using register fun instead of invoke m in New func
		m.statusBar.SetStatus("fetching done")
		return nil
	}, func(err error) error {
		// TODO using register fun instead of invoke m in New func
		m.statusBar.SetStatus("fetching error: " + err.Error())
		return nil
	}))
	// TODO using register fun instead of invoke m in New func
	m.fetcher = fetcher

	// FIXME focus as an method
	m.tasksPane.Focus()

	return &m
}
