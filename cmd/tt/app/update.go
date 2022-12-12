package app

import (
	"fmt"

	"github.com/alswl/go-toodledo/cmd/tt/components"
	comsidebar "github.com/alswl/go-toodledo/cmd/tt/components/sidebar"
	comstatusbar "github.com/alswl/go-toodledo/cmd/tt/components/statusbar"
	"github.com/alswl/go-toodledo/cmd/tt/components/taskspane"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thoas/go-funk"
)

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
		cmd = m.updateFocusedModel(typedMsg)

		// normal keypress

	case RefreshMsg:
		// trigger refresh
		cmd = m.Refresh(typedMsg.isHardRefresh)

	case RefreshTasks:
		// callback of refresh
		tasks, ierr := m.taskLocalSvc.ListAllByQuery(m.states.query)
		if ierr != nil {
			m.log.WithError(ierr).Error("list tasks")
		}
		rts, _ := m.taskRichSvc.RichThem(tasks)
		m.states.Tasks = rts

		_ = m.updateTaskPane(rts)
		m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks reload: %d", len(tasks)))
		// return nil
		return m, cmd

	case tea.WindowSizeMsg:
		// TODO should return m
		m.handleResize(typedMsg)
	}
	return m, cmd
}

// updateFocusedModel updates sub model.
func (m *Model) updateFocusedModel(msg tea.KeyMsg) tea.Cmd {
	var newM tea.Model
	var cmd tea.Cmd
	mm := m.getFocusedModel()

	// post action in main app
	switch typedMM := mm.(type) {
	case *taskspane.Model:
		cmd = m.handleTaskPane(msg)

	case comsidebar.Model:
		m.sidebar, cmd = typedMM.UpdateTyped(msg)
		return cmd

	case comstatusbar.Model:
		newM, cmd = mm.Update(msg)
		m.statusBar, _ = newM.(comstatusbar.Model)
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

// handleCommandMode handles command mode, return false if continued.
func (m *Model) handleCommandMode(msg tea.KeyMsg) (tea.Cmd, bool) {
	switch msg.String() {
	case "tab":
		// change the model fields(isFocused)
		m.loopFocusPane()
		return nil, false
	case "r":
		cmd := m.Refresh(false)
		return cmd, false
	case "R":
		cmd := m.Refresh(true)
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

func (m *Model) handleResize(msg tea.WindowSizeMsg) {
	m.states.width = msg.Width
	m.states.height = msg.Height
	const twoColumns = 2
	const totalColumns = 12
	sideBarWidth := msg.Width * twoColumns / totalColumns
	m.sidebar.Resize(sideBarWidth, msg.Height-1)
	taskPaneWidth := msg.Width - sideBarWidth
	for _, p := range m.tasksPanes {
		p.Resize(taskPaneWidth, msg.Height-1)
	}
	m.statusBar.Resize(msg.Width, 0)
}
func (m *Model) focusStatusBar() {
	m.focus("statusbar")
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
	}
	panic("unreachable")
}

// OnItemChange handle the sidebar menu change.
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

	// m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))
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
