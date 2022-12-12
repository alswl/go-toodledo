package statusbar

import "fmt"

func (m Model) View() string {
	// search mode
	if m.filterTextInput.Focused() {
		m.sb.SetContent("Search", m.filterTextInput.View(), "", "")
		return m.sb.View()
	}

	// display mode
	// append filter indicator TODO move mode to update?
	mode := m.mode
	if m.filterTextInput.Value() != "" {
		mode = fmt.Sprintf("%s /", m.filterTextInput.Value())
	}
	m.sb.SetContent(mode, m.status, m.info1, m.info2)
	return m.sb.View()
}
