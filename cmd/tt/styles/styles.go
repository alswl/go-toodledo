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

	EmptyStyle = lipgloss.NewStyle().Padding(0, 0).Margin(0, 0)

	PaneStyle = EmptyStyle.Copy().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(borderColor)
	// TODO more color
	FocusedPaneStyle = PaneStyle.Copy()
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
