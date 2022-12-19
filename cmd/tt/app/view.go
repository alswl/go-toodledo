package app

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	if m.err != nil {
		m.statusBar.SetMode("ERROR")
		m.statusBar.SetStatus(m.err.Error())
	}

	taskPane := m.getOrCreateTaskPaneByQuery()
	statusBar := m.statusBar.View()

	left := m.sidebar.View()
	var right string
	if m.states.taskDetailID != 0 {
		right = m.taskDetail.View()
	} else {
		right = taskPane.View()
	}
	mainFrame := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		right,
	)
	return styles.EmptyStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			mainFrame,
			statusBar,
		),
	)
}
