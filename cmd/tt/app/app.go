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
	switchAllowedPanes = []string{
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

type RefreshMsg struct {
}

// Model is the main tt app
type Model struct {
	taskRichSvc  services.TaskRichService
	contextSvc   services.ContextService
	folderSvc    services.FolderService
	goalSvc      services.GoalService
	taskSvc      services.TaskExtendedService
	taskLocalSvc services.TaskPersistenceExtService

	// properties
	log logrus.FieldLogger

	// states TODO
	states *States
	err    error
	// focused model: tasks, sidebar, statusbar
	focused string
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
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		sideBarWidth := msg.Width * 2 / 12
		taskPaneWidth := msg.Width - sideBarWidth
		m.sidebar.Resize(sideBarWidth, msg.Height-1)
		m.tasksPane.Resize(taskPaneWidth, msg.Height-1)
		m.statusBar.Resize(msg.Width, 0)
	case RefreshMsg:
		// nothing, only ui refresh
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
				m.loopFocusPane()
				return m, tea.Batch(cmds...)
			case "r":
				err := m.fetcher.Notify(false)
				if err != nil {
					m.log.WithError(err).Error("notify fetcher")
				}
				newCmd := func() tea.Msg {
					select {
					case <-m.fetcher.UIRefresh():
						return RefreshMsg{}
					}
				}
				cmds = append(cmds, newCmd)
				return m, tea.Batch(cmds...)
			case "R":
				err := m.fetcher.Notify(true)
				if err != nil {
					m.log.WithError(err).Error("notify fetcher(force)")
				}
				newCmd := func() tea.Msg {
					select {
					case <-m.fetcher.UIRefresh():
						return RefreshMsg{}
					}
				}
				cmds = append(cmds, newCmd)
				return m, tea.Batch(cmds...)
			}
		}

		// 3. sub focused component
		m, cmd = m.updateFocusedModel(msg)
		cmds = append(cmds, cmd)

		// normal keypress
	}
	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	if m.err != nil {
		m.statusBar.SetMode("ERROR")
		m.statusBar.SetStatus(m.err.Error())
	}

	return styles.EmptyStyle.Render(
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

func (m *Model) loopFocusPane() {
	// only allowed switchAllowedPanes, not support status bar
	if !funk.ContainsString(switchAllowedPanes, m.focused) {
		return
	}

	currentIdx := funk.IndexOf(switchAllowedPanes, m.focused)
	nextPane := switchAllowedPanes[(currentIdx+1)%len(switchAllowedPanes)]

	m.focus(nextPane)
}

func (m *Model) getFocusedModeTyped() components.FocusableInterface {
	name := m.focused
	switch name {
	case "tasks":
		return &m.tasksPane
	case "sidebar":
		return &m.sidebar
	case "statusbar":
		return &m.statusBar
	default:
		panic("unreachable")
	}
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

func (m *Model) focus(next string) {
	m.getFocusedModeTyped().Blur()
	m.focused = next
	m.getFocusedModeTyped().Focus()
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
	tasks, err := m.taskLocalSvc.ListAllByQuery(m.states.query)
	if err != nil {
		m.statusBar.SetStatus("ERROR: " + err.Error())
	}
	rts, _ := m.taskRichSvc.RichThem(tasks)
	m.states.Tasks = rts
	m.tasksPane, _ = m.tasksPane.UpdateTyped(m.states.Tasks)
	//m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))
	m.statusBar.SetInfo1(fmt.Sprintf("./%d", len(m.states.Tasks)))

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
	taskExtSvc := app.TaskExtSvc
	taskLocalSvc := app.TaskExtSvc
	taskRichSvc := app.TaskRichSvc
	contextSvc := app.ContextExtSvc
	folderSvc := app.FolderExtSvc
	goalSvc := app.GoalExtSvc
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
	statusBar.SetInfo1(fmt.Sprintf("./%d", len(states.Tasks)))
	statusBar.SetInfo2("HELP(h)")

	// task pane
	taskPane := taskspane.InitModel(taskExtSvc, states.Tasks)

	sidebar := comsidebar.InitModel(comsidebar.Properties{})

	// main app
	m := Model{
		log:          log,
		taskRichSvc:  taskRichSvc,
		contextSvc:   contextSvc,
		folderSvc:    folderSvc,
		goalSvc:      goalSvc,
		taskSvc:      taskExtSvc,
		taskLocalSvc: taskLocalSvc,
		states:       states,
		err:          nil,
		focused:      "tasks",
		ready:        false,
		tasksPane:    taskPane,
		statusBar:    statusBar,
		sidebar:      sidebar,
		isInputting:  false,
	}

	m.sidebar.RegisterItemChange(m.OnItemChange)

	// init fetcher
	describer := fetchers.NewStatusDescriber(func() error {
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
	})
	fetcher := fetchers.NewSimpleFetcher(log, 1*time.Minute, fetchers.NewToodledoFetchFnPartial(
		log,
		app.FolderExtSvc,
		app.ContextExtSvc,
		app.GoalExtSvc,
		app.TaskExtSvc,
		app.AccountSvc,
	), describer)
	// TODO using register fun instead of invoke m in New func
	m.fetcher = fetcher

	m.tasksPane.Focus()

	return &m
}
