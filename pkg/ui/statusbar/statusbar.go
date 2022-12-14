package statusbar

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Resize(width, _ int) {
	m.sb.SetSize(width)
}

func (m *Model) SetMode(mode string) {
	m.mode = mode
}

func (m *Model) SetStatus(status string) {
	m.status = status
}

func (m *Model) SetInfo1(msg string) {
	m.info1 = msg
}

func (m *Model) SetInfo2(msg string) {
	m.info2 = msg
}

func (m *Model) FocusFilter() {
	m.filterTextInput.Focus()
}

func (m Model) updateFilterTextInput(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	if msgType, ok := msg.(tea.KeyMsg); ok {
		switch msgType.String() {
		case "enter":
			m.filterTextInput.Blur()
		case "esc":
			m.filterTextInput.SetValue("")
			m.filterTextInput.Blur()
		default:
			m.filterTextInput, cmd = m.filterTextInput.Update(msgType)
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

func (m *Model) Info(msg string) {
	m.SetStatus(msg)
}

func (m *Model) Warn(msg string) {
	m.Info(msg)
}

func (m *Model) Error(msg string) {
	m.Info(msg)
}
