package app

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/ui/taskspane"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) handleOpenTaskDetail(pane *taskspane.Model) tea.Cmd {
	id, err := pane.Selected()
	if err != nil {
		m.log.WithField("pane", pane).WithError(err).Warn("get selected task")
		return nil
	}
	m.states.taskDetailID = id

	taskRich, err := m.taskRichSvc.Find(id)
	if err != nil {
		m.log.WithField("id", id).WithError(err).Warn("get task by id")
		return nil
	}

	newM, cmd := m.taskDetail.UpdateTyped(*taskRich)
	m.taskDetail = newM
	return tea.Batch(cmd, func() tea.Msg {
		return models.RefreshUIMsg{}
	})
}
