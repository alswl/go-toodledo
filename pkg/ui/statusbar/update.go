package statusbar

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.filterTextInput.Focused() {
		input, newCmd := m.updateFilterTextInput(msg)
		return input, newCmd
	}

	newM, cmd := m.sb.Update(msg)
	m.sb = newM
	return m, cmd
}
