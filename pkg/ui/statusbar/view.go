package statusbar

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

func (m Model) simpleView() string {
	width := lipgloss.Width

	spinnerView := ""
	if m.loading {
		spinnerView = m.loadingSpinner.View()
	}

	m.loadingSpinner.Style = lipgloss.NewStyle().
		Foreground(m.spinnerColors.Foreground).
		Background(m.spinnerColors.Background).
		Padding(0, 1).
		Height(Height)

	modeC := lipgloss.NewStyle().
		Foreground(m.modeColors.Foreground).
		Background(m.modeColors.Background).
		Padding(0, 1).
		Height(Height).
		Render(truncate.StringWithTail(m.mode, modeWidth, "..."))
	info1C := lipgloss.NewStyle().
		Foreground(m.info1Colors.Foreground).
		Background(m.info1Colors.Background).
		Align(lipgloss.Right).
		Padding(0, 1).
		Height(Height).
		Render(m.info1)
	info2C := lipgloss.NewStyle().
		Foreground(m.info2Colors.Foreground).
		Background(m.info2Colors.Background).
		Padding(0, 1).
		Height(Height).
		Render(m.info2)
	statusC := lipgloss.NewStyle().
		Foreground(m.statusColors.Foreground).
		Background(m.statusColors.Background).
		Padding(0, 1).
		Height(Height).
		Width(m.Width - width(modeC) - width(info1C) - width(info2C) - width(spinnerView)).
		Render(truncate.StringWithTail(
			m.message,
			uint(m.Width-width(modeC)-width(info1C)-width(info2C)-width(spinnerView)-defaultColumPaddingFour),
			"..."),
		)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		modeC,
		statusC,
		spinnerView,
		info1C,
		info2C,
	)
}

func (m Model) View() string {
	// input mode
	if m.input.Focused() {
		m.message = m.input.View()
		m.info1 = ""
		return m.simpleView()
	}

	// display mode
	// append filter indicator TODO move mode to update?
	message := m.message
	if m.input.Value() != "" {
		message = fmt.Sprintf("/%s ", m.input.Value())
	}
	m.Info(message)
	return m.simpleView()
}
