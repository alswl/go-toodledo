package statusbar

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.filterTextInput.Focused() {
		input, newCmd := m.updateFilterTextInput(msg)
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
