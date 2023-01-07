package app

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/alswl/go-toodledo/pkg/ui/detail"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/ui"
	"github.com/alswl/go-toodledo/pkg/ui/sidebar"
	comstatusbar "github.com/alswl/go-toodledo/pkg/ui/statusbar"
	"github.com/alswl/go-toodledo/pkg/ui/taskspane"
	"github.com/alswl/go-toodledo/pkg/utils"

	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thoas/go-funk"
)

const lastQueryKey = "last-query"
const sidebarStatesKey = "sidebar-states"

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// FIXME do not using pointer

	// process logics
	// 1. global keymap
	// 2. input TODO input is one of focused component?
	// 3. focused component

	var cmd tea.Cmd
	switch typedMsg := msg.(type) {
	case tea.KeyMsg: // handle event bubble
		// 1. global keymap
		if funk.ContainsString([]string{"ctrl+c"}, typedMsg.String()) {
			if typedMsg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		}

		// 2. main app, command mode
		if !m.isInputting {
			newCmd, isContinue := m.handleCommandMode(typedMsg)
			if !isContinue {
				return m, newCmd
			}
		}

		// 3. sub focused component
		cmd = m.keyPressFocusedModel(typedMsg)

	case models.FetchTasksMsg:
		// trigger refresh
		cmd = m.FetchTasks(typedMsg.IsHardRefresh)

	case models.RefreshPropertiesMsg:
		cmd = m.ReloadDependencies()

	case models.RefreshTasksMsg:
		// refresh tasks(ui)
		cmd = m.ReloadTasks()
		return m, cmd

	case models.ReturnMsg:
		// return from sub component
		m.states.taskDetailID = 0
		m.focus("tasks")

	case tea.WindowSizeMsg:
		cmd = m.handleResize(typedMsg)
	default:
		// all others message broadcast to sub component
		m.statusBar, cmd = m.statusBar.UpdateTyped(msg)
		// TODO pane task and sidebar
	}
	return m, cmd
}

// keyPressFocusedModel updates sub model.
func (m *Model) keyPressFocusedModel(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	mm := m.getFocusedModel()

	// post action in main app
	switch typedMM := mm.(type) {
	case *taskspane.Model:
		cmd = m.handleTaskPane(msg)

	case sidebar.Model:
		m.sidebar, cmd = typedMM.UpdateTyped(msg)
		return cmd

	case comstatusbar.Model:
		var newM comstatusbar.Model
		newM, cmd = typedMM.UpdateTyped(msg)
		m.statusBar = newM
		m.getOrCreateTaskPaneByQuery().Filter(m.statusBar.GetFilterInput())

		// post action
		switch msg.String() {
		case "enter", "esc":
			m.isInputting = false
			m.focus("tasks")
		}
		return cmd
	case detail.Model:
		var newM detail.Model
		newM, cmd = typedMM.UpdateTyped(msg)
		m.taskDetail = newM
	}
	return cmd
}

// handleCommandMode handles command mode, return false if continued.
func (m *Model) handleCommandMode(msg tea.KeyMsg) (tea.Cmd, bool) {
	switch msg.String() {
	case "tab":
		// change the model fields(isFocused)
		m.loopFocusPane()
		return nil, false
	case "r":
		cmd := m.FetchTasks(false)
		return cmd, false
	case "R":
		cmd := m.FetchTasks(true)
		return cmd, false
	}
	return nil, true
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

// handleResize handles window resize event.
// once the app started, it will be called with msg.
func (m *Model) handleResize(msg tea.WindowSizeMsg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.Width = msg.Width
	m.Height = msg.Height
	const twoColumns = 2
	const totalColumns = 12
	sideBarWidth := msg.Width * twoColumns / totalColumns
	mainPaneWidth := msg.Width - sideBarWidth

	m.sidebar.Resize(sideBarWidth, msg.Height-1)
	m.sidebar, cmd = m.sidebar.UpdateTyped(msg)
	cmds = append(cmds, cmd)

	for _, p := range m.tasksPanes {
		p.Resize(mainPaneWidth, msg.Height-1)
		// FIXME resize msg will cause wrong height calculation
		// find a way to fix it, and enable this msg
		// var newCmd tea.Cmd
		// newPane, newCmd := p.UpdateTyped(msg)
		// m.tasksPanes[key] = &newPane
		// cmds = append(cmds, newCmd)
	}
	m.statusBar.Resize(msg.Width, 0)
	m.statusBar, cmd = m.statusBar.UpdateTyped(msg)
	cmds = append(cmds, cmd)

	m.taskDetail.Resize(mainPaneWidth, msg.Height)
	m.taskDetail, cmd = m.taskDetail.UpdateTyped(msg)

	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

func (m *Model) focusStatusBar() {
	m.focus("statusbar")
}

func (m *Model) getFocusedModeTyped() ui.FocusableInterface {
	name := m.focused
	switch name {
	case "tasks":
		return m.getOrCreateTaskPaneByQuery()
	case "sidebar":
		return &m.sidebar
	case "statusbar":
		return &m.statusBar
	case "detail":
		return &m.taskDetail
	default:
		panic("unreachable")
	}
}

func (m *Model) focus(next string) {
	m.getFocusedModeTyped().Blur()
	m.focused = next
	m.getFocusedModeTyped().Focus()
}

func (m *Model) getFocusedModel() tea.Model {
	switch m.focused {
	case "tasks":
		return m.getOrCreateTaskPaneByQuery()
	case "sidebar":
		return m.sidebar
	case "statusbar":
		return m.statusBar
	case "detail":
		return m.taskDetail
	}
	panic("unreachable")
}

// OnItemChange handle the sidebar menu change.
func (m *Model) OnItemChange(tab string, item sidebar.Item) error {
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

	// m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))
	m.statusBar.SetInfo1(fmt.Sprintf("./%d", len(m.states.Tasks)))

	// save sidebar for restore
	bytes, _ := yaml.Marshal(m.states.query)
	err = m.settingSvc.Put(lastQueryKey, string(bytes))
	if err != nil {
		m.log.WithError(err).Error("put last query failed")
	}
	states := m.sidebar.GetStates()
	bytes, _ = yaml.Marshal(states)
	err = m.settingSvc.Put(sidebarStatesKey, string(bytes))
	if err != nil {
		m.log.WithError(err).Error("put sidebar states failed")
	}

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

func (m *Model) ReloadDependencies() tea.Cmd {
	cs, err := m.contextExtSvc.ListAll()
	if err != nil {
		m.err = err
		return nil
	}
	m.states.Contexts = cs
	// Contexts are first tab in sidebar
	m.states.Contexts = append([]*models.Context{{
		ID:   0,
		Name: "All",
	}}, cs...)
	m.states.Contexts = append(m.states.Contexts, &models.Context{
		ID:   -1,
		Name: "None",
	})
	fs, err := m.folderExtSvc.ListAll()
	if err != nil {
		m.err = err
		return nil
	}
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

	gs, err := m.goalExtSvc.ListAll()
	if err != nil {
		m.err = err
	}
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

	m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Contexts))
	m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Folders))
	m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Goals))

	return nil
}

// ReloadTasks refresh local ui.
func (m *Model) ReloadTasks() tea.Cmd {
	tasks, err := m.taskExtSvc.ListAllByQuery(m.states.query)
	if err != nil {
		m.statusBar.SetStatus("ERROR: " + err.Error())
	}
	rts, _ := m.taskRichSvc.RichThem(tasks)
	m.states.Tasks = rts
	cmd := m.updateTaskPane(rts)
	return cmd
}
