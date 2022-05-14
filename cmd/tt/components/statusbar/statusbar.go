package statusbar

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/tt/components"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/teacup/statusbar"
)

type Model struct {
	sb statusbar.Bubble
	components.Focusable
	filterTextInput textinput.Model

	//item *models.RichTask

	// in statusBar
	mainContent string
}

func (m *Model) Resize(width, _ int) {
	m.sb.SetSize(width)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.filterTextInput.Focused() {
		input, cmd := m.updateFilterTextInput(msg)
		return input, cmd
	}

	newM, cmd := m.sb.Update(msg)
	m.sb = newM
	return m, cmd
}

func (m Model) View() string {
	// search mode
	if m.filterTextInput.Focused() {
		m.sb.SetContent("Search", m.filterTextInput.View(), "", "")
		return m.sb.View()
	}

	// append filter indicator
	first := m.sb.FirstColumn
	if m.filterTextInput.Value() != "" {
		first = fmt.Sprintf("%s /", m.sb.FirstColumn)
	}
	m.sb.SetContent(first, m.sb.SecondColumn, m.sb.ThirdColumn, m.sb.FourthColumn)
	return m.sb.View()
}

func (m Model) genSecondColumn() string {
	second := m.mainContent
	if m.filterTextInput.Value() != "" {
		second = fmt.Sprintf("%s%s %s", m.filterTextInput.Prompt, m.filterTextInput.Value(), m.sb.SecondColumn)
	}
	return second
}

func (m *Model) SetContent(firstColumn, secondColumn, thirdColumn, fourthColumn string) {
	m.sb.SetContent(firstColumn, secondColumn, thirdColumn, fourthColumn)
}

func (m *Model) FocusFilter() {
	m.filterTextInput.Focus()
}

func (m Model) updateFilterTextInput(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.filterTextInput.Blur()
		case "esc":
			m.filterTextInput.SetValue("")
			m.filterTextInput.Blur()
		default:
			m.filterTextInput, cmd = m.filterTextInput.Update(msg)
		}
	}

	return m, cmd
}

func (m Model) GetFilterText() string {
	return m.filterTextInput.Value()
}

func (m Model) GetFilterInput() textinput.Model {
	return m.filterTextInput
}

func NewDefault() Model {
	sb := statusbar.New(
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.Pink, Dark: styles.Pink},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkGray, Dark: styles.DarkGray},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkPurple, Dark: styles.DarkPurple},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkPurple, Dark: styles.DarkPurple},
		},
	)
	ti := textinput.New()
	ti.Prompt = "/"
	return Model{sb: sb, filterTextInput: ti}
}
