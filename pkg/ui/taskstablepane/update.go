package taskstablepane

import (
	"fmt"

	"github.com/alswl/go-toodledo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)

	switch msgType := msg.(type) {
	case []*models.RichTask:
		// update tasks(render new table)
		m.tableModel = m.tableModel.WithRows(TasksRenderRows(msgType))

	case tea.WindowSizeMsg:
		// top, right, bottom, left := docStyle.GetMargin()
		m.Resize(msgType.Width, msgType.Height)

	default:
		m.tableModel, cmd = m.tableModel.Update(msg)
	}

	// post event
	current, total := m.CurrentAndTotalPage()
	cmd = tea.Batch(cmd, func() tea.Msg {
		return models.StatusMsg{Info2: fmt.Sprintf("%d/%d", current, total)}
	})

	return m, cmd
}

func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}
