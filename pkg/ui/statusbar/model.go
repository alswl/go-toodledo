package statusbar

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/ui"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Height represents the height of the statusbar.
const Height = 1
const modeWidth = 30
const defaultColumPaddingFour = 3

// Model is the statusbar model.
// it was inspired by teacup https://github.com/knipferrc/teacup/blob/main/statusbar/statusbar.go
type Model struct {
	ui.Focusable
	ui.Resizable

	// states
	mode    string
	status  string
	loading bool
	info1   string
	info2   string

	modeColors    common.ColorConfig
	statusColors  common.ColorConfig
	spinnerColors common.ColorConfig
	info1Colors   common.ColorConfig
	info2Colors   common.ColorConfig

	// view
	spinner         spinner.Model
	filterTextInput textinput.Model
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func NewDefault() Model {
	ti := textinput.New()
	ti.Prompt = "/"
	s := spinner.New()
	s.Spinner = spinner.Dot

	m := Model{
		modeColors: common.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.Pink, Dark: styles.Pink},
		},
		statusColors: common.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkGray, Dark: styles.DarkGray},
		},
		spinnerColors: common.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.Pink, Dark: styles.Pink},
		},
		info1Colors: common.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkPurple, Dark: styles.DarkPurple},
		},
		info2Colors: common.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkPurple, Dark: styles.DarkPurple},
		},
		filterTextInput: ti,
		spinner:         s,
	}
	return m
}
