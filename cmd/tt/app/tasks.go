package app

import (
	"fmt"

	"github.com/alswl/go-toodledo/pkg/models"

	"github.com/alswl/go-toodledo/pkg/ui/taskspane"

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
	newP := taskspane.InitModel(m.states.Tasks, m.states.width, m.states.height-1)
	m.tasksPanes[key] = &newP
	return &newP
}

func (m *Model) FetchTasks(isHardRefresh bool) tea.Cmd {
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
		return models.RefreshTasksMsg{}
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
		// TODO inline
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
		case "o":
			// open task detail pane
			m.focus("detail")
			cmd = m.handleOpenTask(pane)
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
		return models.FetchTasksMsg{IsHardRefresh: false}
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
		return models.FetchTasksMsg{IsHardRefresh: false}
	}
}
