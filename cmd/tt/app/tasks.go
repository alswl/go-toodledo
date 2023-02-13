package app

import (
	"fmt"

	"github.com/alswl/go-toodledo/pkg/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) FetchTasks(isHardRefresh bool) tea.Cmd {
	// detect refresh in progress
	ok := refreshLock.TryLock()
	if !ok {
		m.log.Warn("refresh in progress")
		return nil
	}
	defer refreshLock.Unlock()
	return tea.Batch(
		func() tea.Msg {
			return m.statusBar.StartSpinner()
		},
		// io blocking in msg
		func() tea.Msg {
			err := m.fetcher.Fetch(isHardRefresh)
			m.log.WithField("hard", isHardRefresh).Info("refreshing done")
			if err != nil {
				m.log.WithError(err).Error("refresh failed")
				m.statusBar.Error(fmt.Sprintf("refresh failed, %s", err.Error()))
				return nil
			}
			m.statusBar.StopSpinner()
			return models.RefreshTasksMsg{}
		},
	)
}

func (m *Model) handleCompleteToggle(id int64) tea.Cmd {
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

func (m *Model) handleTimerToggle(id int64) tea.Cmd {
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
