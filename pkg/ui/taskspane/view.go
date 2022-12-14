package taskspane

import "github.com/alswl/go-toodledo/cmd/tt/styles"

func (m Model) View() string {
	m.Viewport.SetContent(
		m.tableModel.View(),
	)

	style := styles.PaneStyle.Copy()
	if m.IsFocused() {
		style = styles.FocusedPaneStyle.Copy()
	}
	return style.Render(m.Viewport.View())
}
