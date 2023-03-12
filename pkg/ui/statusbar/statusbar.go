package statusbar

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/models"
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
const ModeSearch = "Search"
const ModeNew = "New"
const ModeDefault = "Default"

// Model is the statusbar model.
// it was inspired by teacup https://github.com/knipferrc/teacup/blob/main/statusbar/statusbar.go
type Model struct {
	ui.Focusable
	ui.Resizable

	// states
	mode    string
	message string
	info1   string
	info2   string

	loading bool

	modeColors    common.ColorConfig
	statusColors  common.ColorConfig
	spinnerColors common.ColorConfig
	info1Colors   common.ColorConfig
	info2Colors   common.ColorConfig

	// view
	loadingSpinner spinner.Model
	input          textinput.Model
}

func (m Model) Init() tea.Cmd {
	return nil
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
		input:          ti,
		loadingSpinner: s,
	}
	return m
}

func (m *Model) Resize(width, _ int) {
	m.Width = width
}

func (m *Model) SetMode(mode string) {
	m.mode = mode
}

func (m Model) GetMode() string {
	return m.mode
}

func (m *Model) SetMessage(message models.StatusMsg) {
	if message.ClearMode {
		m.mode = ""
	} else if message.Mode != "" {
		m.mode = message.Mode
	}
	if message.ClearMessage {
		m.message = ""
	} else if message.Message != "" {
		m.message = message.Message
	}
	if message.ClearInfo1 {
		m.info1 = ""
	} else if message.Info1 != "" {
		m.info1 = message.Info1
	}
	if message.ClearInfo2 {
		m.info2 = ""
	} else if message.Info2 != "" {
		m.info2 = message.Info2
	}
}

func (m *Model) SetInfo1(msg string) {
	m.info1 = msg
}

func (m *Model) SetInfo2(msg string) {
	m.info2 = msg
}

func (m Model) GetInputText() string {
	return m.input.Value()
}

func (m Model) GetInput() textinput.Model {
	return m.input
}

func (m *Model) Info(msg string) {
	m.SetMessage(models.StatusMsg{
		Message: msg,
	})
}

func (m *Model) Warn(msg string) {
	m.Info("Warn: " + msg)
}

func (m *Model) Error(msg string) {
	m.Info("Error: " + msg)
}
