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
	"strconv"
	"sync"
	"time"
)

var (
	switchAllowedPanes = []string{
		"tasks",
		"sidebar",
	}
)

type States struct {
	width  int
	height int

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
// it was singleton
type Model struct {
	taskRichSvc  services.TaskRichService
	contextSvc   services.ContextService
	folderSvc    services.FolderService
	goalSvc      services.GoalService
	taskExtSvc   services.TaskExtendedService
	taskLocalSvc services.TaskPersistenceExtService

	// properties
	log         logrus.FieldLogger
	refreshLock sync.Mutex

	// states TODO
	states *States
	err    error
	// focused model: tasks, sidebar, statusbar
	focused string
	// TODO ready check
	ready bool
	//isSidebarOpen bool

	// view
	tasksPanes map[string]*taskspane.Model
	sidebar    comsidebar.Model
	statusBar  comstatusbar.Model
	// TODO help pane
	//help          help.Model
	isInputting bool
	fetcher     fetchers.DaemonFetcher
}

func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	// states init
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

		// TODO using last selected menu

		// tasks
		tasks, err := m.taskExtSvc.ListAllByQuery(m.states.query)
		if err != nil {
			m.statusBar.SetStatus("ERROR: " + err.Error())
		}
		rts, _ := m.taskRichSvc.RichThem(tasks)
		m.states.Tasks = rts

		cmd := m.updateTaskPaneByQuery(rts)
		cmds = append(cmds, cmd)
		m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

		return nil
	})

	// daemon fetcher sstart
	cmds = append(cmds, func() tea.Msg {
		m.fetcher.Start(context.Background())
		return nil
	})

	// refresh at start
	cmds = append(cmds, func() tea.Msg {
		return m.Refresh(false)
	})

	return tea.Batch(cmds...)
}

// handleCommandMode handles command mode, return false if continue
func (m *Model) handleCommandMode(msg tea.KeyMsg) (tea.Cmd, bool) {
	switch msg.String() {
	case "tab":
		// change the model fields(isFocused)
		m.loopFocusPane()
		return nil, false
	case "r":
		cmd := m.handleRefresh(false)
		return cmd, false
	case "R":
		cmd := m.handleRefresh(true)
		return cmd, false
	}
	return nil, true
}

func (m *Model) Refresh(isHardRefresh bool) tea.Cmd {
	return m.handleRefresh(isHardRefresh)
}

func (m *Model) handleRefresh(isHardRefresh bool) tea.Cmd {
	// detect refresh in progress
	ok := m.refreshLock.TryLock()
	if !ok {
		return nil
	}
	defer m.refreshLock.Unlock()

	refreshedChan, err := m.fetcher.Notify(isHardRefresh)
	if err != nil {
		m.log.WithError(err).Error("notify fetcher, hard(?)" + strconv.FormatBool(isHardRefresh))
	}
	// this cmd is works like promise
	cmd := func() tea.Msg {
		select {
		case success := <-refreshedChan:
			if !success {
				m.statusBar.SetStatus("ERROR: refresh failed")
				return nil
			}
			tasks, ierr := m.taskLocalSvc.ListAllByQuery(m.states.query)
			if ierr != nil {
				m.log.WithError(ierr).Error("list tasks")
			}
			rts, _ := m.taskRichSvc.RichThem(tasks)
			m.states.Tasks = rts

			_ = m.updateTaskPaneByQuery(rts)
			return RefreshMsg{}
		}
	}
	return cmd
}

func (m *Model) handleResize(msg tea.WindowSizeMsg) {
	m.states.width = msg.Width
	m.states.height = msg.Height
	sideBarWidth := msg.Width * 2 / 12
	m.sidebar.Resize(sideBarWidth, msg.Height-1)
	taskPaneWidth := msg.Width - sideBarWidth
	for _, p := range m.tasksPanes {
		p.Resize(taskPaneWidth, msg.Height-1)
	}
	m.statusBar.Resize(msg.Width, 0)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// process logics
	// 1. global keymap
	// 2. input TODO input is one of focused component?
	// 3. focused component

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleResize(msg)
	case RefreshMsg:
		// nothing, only ui refresh
		return m, cmd
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
			newCmd, isContinue := m.handleCommandMode(msg)
			if !isContinue {
				return m, newCmd
			}
		}

		// 3. sub focused component
		cmd = m.updateFocusedModel(msg)

		// normal keypress
	}
	return m, cmd
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
				m.getOrCreateTaskPaneByQuery().View(),
			),
			m.statusBar.View(),
		),
	)
}

func (m *Model) getOrCreateTaskPaneByQuery() *taskspane.Model {
	key := ""
	if m.states.query != nil {
		key = m.states.query.UniqString()
	}
	if p, ok := m.tasksPanes[key]; ok {
		return p
	} else {
		newP := taskspane.InitModel(m.taskExtSvc, m.states.Tasks, m)
		//// trigger ui redraw
		m.tasksPanes[key] = &newP
		m.handleResize(tea.WindowSizeMsg{
			Width:  m.states.width,
			Height: m.states.height,
		})
		return &newP
	}
}

func (m *Model) updateTaskPaneByQuery(msg tea.Msg) tea.Cmd {
	key := ""
	if m.states.query != nil {
		key = m.states.query.UniqString()
	}
	pane := m.getOrCreateTaskPaneByQuery()
	newM, newMsg := pane.UpdateTyped(msg)
	m.tasksPanes[key] = &newM
	return newMsg
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
		return m.getOrCreateTaskPaneByQuery()
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
		return m.getOrCreateTaskPaneByQuery()
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

func (m *Model) handleTaskPane(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	switch msg.String() {
	case "/":
		m.isInputting = true
		m.focusStatusBar()
		m.statusBar.FocusFilter()
	default:
		cmd = m.updateTaskPaneByQuery(msg)
	}
	return cmd
}

func (m *Model) updateFocusedModel(msg tea.KeyMsg) tea.Cmd {
	var newM tea.Model
	var cmd tea.Cmd
	mm := m.getFocusedModel()

	// post action in main app
	switch mm.(type) {
	case *taskspane.Model:
		cmd = m.handleTaskPane(msg)
	case comsidebar.Model:
		mmTyped := mm.(comsidebar.Model)
		m.sidebar, cmd = mmTyped.UpdateTyped(msg)
		return cmd
	case comstatusbar.Model:
		newM, cmd = mm.Update(msg)
		m.statusBar = newM.(comstatusbar.Model)
		m.getOrCreateTaskPaneByQuery().Filter(m.statusBar.GetFilterInput())

		// post action
		switch msg.String() {
		case "enter", "esc":
			m.isInputting = false
			m.focus("tasks")
		}
		return cmd
	}
	return cmd
}

// OnItemChange handle the sidebar menu change
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

	//m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))
	m.statusBar.SetInfo1(fmt.Sprintf("./%d", len(m.states.Tasks)))

	return nil
}

func (m *Model) Info(msg string) {
	m.statusBar.Info(msg)
}

func (m *Model) Warn(msg string) {
	m.statusBar.Warn(msg)
}

func (m *Model) Error(msg string) {
	m.statusBar.Error(msg)
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
	sidebar := comsidebar.InitModel(comsidebar.Properties{})

	// main app
	m := Model{
		log:          log,
		taskRichSvc:  taskRichSvc,
		contextSvc:   contextSvc,
		folderSvc:    folderSvc,
		goalSvc:      goalSvc,
		taskExtSvc:   taskExtSvc,
		taskLocalSvc: taskLocalSvc,
		states:       states,
		err:          nil,
		focused:      "tasks",
		ready:        false,
		statusBar:    statusBar,
		sidebar:      sidebar,
		isInputting:  false,
		tasksPanes:   map[string]*taskspane.Model{},
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

	m.getOrCreateTaskPaneByQuery().Focus()

	return &m
}
