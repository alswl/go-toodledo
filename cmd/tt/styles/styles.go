package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var (
	// color {{{
	// from gh-dash
	// my.
	borderColor = lipgloss.AdaptiveColor{Light: "#212F3D", Dark: "#D5D8DC"}
	// }}}.

	NoStyle = lipgloss.NewStyle().Padding(0, 0).Margin(0, 0)

	PaneStyle = NoStyle.Copy().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(borderColor)
	FocusedPaneStyle = PaneStyle.Copy().BorderForeground(lipgloss.AdaptiveColor{Dark: "#F25D94", Light: "#F25D94"})
	EmptyBorderStyle = lipgloss.Border{
		Top:         "",
		Bottom:      "",
		Left:        "",
		Right:       "",
		TopLeft:     "",
		TopRight:    "",
		BottomRight: "",
		BottomLeft:  "",
	}
	EmptyTableBorderStyle = table.Border{
		Top:         "",
		Bottom:      "",
		Left:        "",
		Right:       "",
		TopLeft:     "",
		TopRight:    "",
		BottomRight: "",
		BottomLeft:  "",
	}
	// }}}.
)
