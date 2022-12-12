package app

import (
	"fmt"

	"github.com/alswl/go-toodledo/cmd/tt/components/taskspane"
	tea "github.com/charmbracelet/bubbletea"
)

// getOrCreateTaskPaneByQuery returns the task pane by query.
// app holds a map of task panes, and the key is the query string.
func (m *Model) getOrCreateTaskPaneByQuery() *taskspane.Model {
	key := ""
	if m.states.query != nil {
		key = m.states.query.UniqString()
	}
	if p, ok := m.tasksPanes[key]; ok {
		return p
	}
	newP := taskspane.InitModel(m.states.Tasks, m)
	// trigger ui redraw
	m.tasksPanes[key] = &newP
	m.handleResize(tea.WindowSizeMsg{
		Width:  m.states.width,
		Height: m.states.height,
	})
	return &newP
}

func (m *Model) Refresh(isHardRefresh bool) tea.Cmd {
	// detect refresh in progress
	ok := refreshLock.TryLock()
	if !ok {
		m.log.Warn("refresh in progress")
		return nil
	}
	defer refreshLock.Unlock()

	// io blocking in msg
	return func() tea.Msg {
		err := m.fetcher.Fetch(isHardRefresh)
		m.log.WithField("hard", isHardRefresh).Info("refreshing done")
		if err != nil {
			m.log.WithError(err).Error("refresh failed")
			m.statusBar.SetStatus(fmt.Sprintf("ERROR: refresh failed, %s", err.Error()))
			return nil
		}
		return RefreshTasks{}
	}
}

func (m *Model) handleTaskPane(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	switch msg.String() {
	case "/":
		m.isInputting = true
		m.focusStatusBar()
		m.statusBar.FocusFilter()
	default:
		cmd = m.updateTaskPane(msg)
	}
	return cmd
}

func (m *Model) updateTaskPane(msg tea.Msg) tea.Cmd {
	paneKey := ""
	if m.states.query != nil {
		paneKey = m.states.query.UniqString()
	}
	pane := m.getOrCreateTaskPaneByQuery()

	var cmd tea.Cmd
	var newM taskspane.Model
	switch msgType := msg.(type) {
	case tea.KeyMsg:
		switch msgType.String() {
		case "x":
			cmd = m.handleCompleteToggle(pane)
		case "enter":
			cmd = m.handleTimerToggle(pane)
		default:
			newM, cmd = pane.UpdateTyped(msg)
			m.tasksPanes[paneKey] = &newM
		}
	default:
		newM, cmd = pane.UpdateTyped(msg)
		m.tasksPanes[paneKey] = &newM
	}

	return cmd
}

func (m *Model) handleCompleteToggle(pane *taskspane.Model) tea.Cmd {
	id, err := pane.Selected()
	if err != nil {
		m.log.WithError(err).Warn("get selected task")
	}
	task, err := m.taskLocalSvc.FindByID(id)
	if err != nil {
		return nil
	}

	return func() tea.Msg {
		if task.Completed == 0 {
			_, ierr := m.taskLocalSvc.Complete(id)
			if ierr != nil {
				m.Error(ierr.Error())
				return nil
			}
		} else {
			_, ierr := m.taskLocalSvc.UnComplete(id)
			if ierr != nil {
				m.Error(ierr.Error())
				return nil
			}
		}
		return RefreshMsg{false}
	}
}

func (m *Model) handleTimerToggle(pane *taskspane.Model) tea.Cmd {
	id, err := pane.Selected()
	if err != nil {
		m.log.WithError(err).Warn("get selected task")
	}
	t, err := m.taskLocalSvc.FindByID(id)
	if err != nil {
		return nil
	}
	if err != nil {
		m.Error(err.Error())
		return nil
	}
	if t.Completed > 0 {
		m.Warn("can't start timer on completed task")
		return nil
	}

	return func() tea.Msg {
		if t.Timeron == 0 {
			// start timer
			err = m.taskLocalSvc.Start(id)
			if err != nil {
				m.Error(err.Error())
				return nil
			}
		} else {
			// stop timer
			err = m.taskLocalSvc.Stop(id)
			if err != nil {
				m.Error(err.Error())
				return nil
			}
		}
		return RefreshMsg{false}
	}
}
