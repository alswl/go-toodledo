package taskstablepane

import "github.com/alswl/go-toodledo/cmd/tt/styles"

func (m Model) View() string {
	m.Viewport.SetContent(
		m.tableModel.View(),
	)
	style := styles.NoStyle.Copy()
	return style.Render(m.Viewport.View())
}
