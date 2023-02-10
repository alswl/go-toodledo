package statusbar

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.input.Focused() {
		input, newCmd := m.updateTextInput(msg)
		return input, newCmd
	}

	switch typedMsg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Resize(typedMsg.Width, typedMsg.Height)
	default:
		m.spinner, cmd = m.spinner.Update(msg)
	}

	return m, cmd
}

func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

func (m *Model) StartSpinner() {
	m.loading = true
	m.spinner.Tick()
}

func (m *Model) StopSpinner() {
	m.loading = false
	newSpinner := spinner.New()
	newSpinner.Style = m.spinner.Style
	m.spinner = newSpinner
}

func (m Model) updateTextInput(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	if msgType, ok := msg.(tea.KeyMsg); ok {
		switch msgType.String() {
		case "enter":
			if m.mode == ModeNew {
				// exit mode new
				m.mode = ModeDefault
				m.message = ""
				m.input.SetValue("")
			}
			m.input.Blur()
		case "esc":
			// clear and cancel
			m.message = ""
			m.input.SetValue("")
			m.input.Blur()
		default:
			m.input, cmd = m.input.Update(msgType)
		}
	}

	return m, cmd
}
func (m *Model) FocusInputSearch() {
	m.mode = ModeSearch
	m.input.Focus()
}

func (m *Model) FocusInputNew() {
	m.mode = ModeNew
	m.input.Focus()
}
