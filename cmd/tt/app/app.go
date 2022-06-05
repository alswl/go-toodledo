package app

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/cmd/tt/components"
	comsidebar "github.com/alswl/go-toodledo/cmd/tt/components/sidebar"
	comstatusbar "github.com/alswl/go-toodledo/cmd/tt/components/statusbar"
	"github.com/alswl/go-toodledo/cmd/tt/components/taskspane"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer"
	"github.com/alswl/go-toodledo/pkg/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thoas/go-funk"
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

type Model struct {
	taskRichSvc services.TaskRichService
	contextSvc  services.ContextCachedService
	folderSvc   services.FolderCachedService
	goalSvc     services.GoalCachedService
	taskSvc     services.TaskCachedService

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
	syncer      syncer.ToodledoFetcher
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
		m.sidebar, _ = m.sidebar.UpdateX(utils.UnwrapListPointer(cs))
		if len(m.states.Contexts) > 0 {
			// using default first now TODO add non-context item
			m.states.query.ContextID = m.states.Contexts[0].ID
		}

		// folders
		m.states.Folders = fs
		m.sidebar, _ = m.sidebar.UpdateX(utils.UnwrapListPointer(fs))

		// goals
		m.states.Goals = gs
		m.sidebar, _ = m.sidebar.UpdateX(utils.UnwrapListPointer(gs))

		// tasks
		tasks, err := m.taskSvc.ListAllByQuery(m.states.query)
		if err != nil {
			m.statusBar.SetStatus("ERROR: " + err.Error())
		}
		rts, _ := m.taskRichSvc.RichThem(tasks)
		m.states.Tasks = rts
		m.tasksPane, _ = m.tasksPane.UpdateX(m.states.Tasks)
		m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

		return nil
	})
	cmds = append(cmds, func() tea.Msg {
		m.statusBar.SetInfo1("syncing")
		err := m.syncer.SyncOnce()
		if err != nil {
			m.statusBar.SetStatus("ERROR: " + err.Error())
		}
		m.statusBar.SetInfo1("done")
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
	switch tab {
	case constants.Contexts:
		m.states.query = &queries.TaskListQuery{}
		if item.ID() == 0 && len(m.states.Contexts) > 0 {
			if len(m.states.Contexts) == 0 {
				// TODO None folder impl
			} else {
				m.states.query.ContextID = m.states.Contexts[0].ID
			}
		} else {
			m.states.query.ContextID = item.ID()
		}
	case constants.Folders:
		m.states.query = &queries.TaskListQuery{}
		if item.ID() == 0 {
			if len(m.states.Folders) == 0 {
				// TODO None folder impl
			} else {
				m.states.query.FolderID = m.states.Folders[0].ID
			}
		} else {
			m.states.query.FolderID = item.ID()
		}
	case constants.Goals:
		m.states.query = &queries.TaskListQuery{}
		if item.ID() == 0 {
			if len(m.states.Goals) == 0 {
				// TODO None folder impl
			} else {
				m.states.query.GoalID = m.states.Goals[0].ID
			}
		} else {
			m.states.query.GoalID = item.ID()
		}
	}
	tasks, err := svc.ListAllByQuery(m.states.query)
	if err != nil {
		m.statusBar.SetStatus("ERROR: " + err.Error())
	}
	rts, _ := taskRichSvc.RichThem(tasks)
	m.states.Tasks = rts
	m.tasksPane, _ = m.tasksPane.UpdateX(m.states.Tasks)
	m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

	return nil
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
	taskSvc, err := injector.InitTaskCachedService()
	if err != nil {
		panic(err)
	}
	taskRichSvc, err := injector.InitTaskRichService()
	if err != nil {
		panic(err)
	}
	contextSvc, err := injector.InitContextCachedService()
	if err != nil {
		panic(err)
	}
	folderSvc, err := injector.InitFolderCachedService()
	if err != nil {
		panic(err)
	}
	goalSvc, err := injector.InitGoalCachedService()
	if err != nil {
		panic(err)
	}
	syncer, err := injector.InitSyncer()
	if err != nil {
		panic(err)
	}

	states := &States{
		Tasks:    []*models.RichTask{},
		Contexts: []*models.Context{},
		Folders:  []*models.Folder{},
		Goals:    []*models.Goal{},
		query:    &queries.TaskListQuery{},
	}

	statusBar := comstatusbar.NewDefault()
	statusBar.SetMode("tasks")
	statusBar.SetStatus("a fox jumped over the lazy dog")
	statusBar.SetInfo1("1/999")
	statusBar.SetInfo2("HELP")

	// FIXME tasks should comes from syncer
	taskPane := taskspane.InitModel(states.Tasks)

	m := Model{
		taskRichSvc: taskRichSvc,
		contextSvc:  contextSvc,
		folderSvc:   folderSvc,
		goalSvc:     goalSvc,
		taskSvc:     taskSvc,
		syncer:      syncer,
		states:      states,
		err:         nil,
		focused:     "tasks",
		tabIndex:    0,
		ready:       false,
		tasksPane:   taskPane,
		statusBar:   statusBar,
		isInputting: false,
	}
	m.sidebar = comsidebar.InitModel(comsidebar.Properties{}, m.OnItemChange)
	// FIXME focus as an method
	m.tasksPane.Focus()

	return &m
}
