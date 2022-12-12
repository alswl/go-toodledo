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
