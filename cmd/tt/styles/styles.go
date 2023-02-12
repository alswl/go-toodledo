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

	// SuccessStyle and others system styles comes from Ant Design
	// https://ant-design.antgroup.com/docs/spec/colors-cn#%E5%8A%9F%E8%83%BD%E8%89%B2
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#57C22D")).Bold(true)
	InfoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#217BFB")).Bold(true)
	WarningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F8AC30")).Bold(true)
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FD4F54")).Bold(true)

	ProcessingStyle = SuccessStyle
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
