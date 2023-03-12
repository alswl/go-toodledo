package app

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	statusBar := m.statusBar.View()

	left := m.sidebar.View()
	right := m.primaryPane.View()
	var mainFrame string
	if m.states.isSidebarVisible {
		mainFrame = lipgloss.JoinHorizontal(
			lipgloss.Top,
			left,
			right,
		)
	} else {
		mainFrame = right
	}
	return styles.NoStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			mainFrame,
			statusBar,
		),
	)
}
