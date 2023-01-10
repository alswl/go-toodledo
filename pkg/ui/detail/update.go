package detail

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// nolint
	switch msgTyped := msg.(type) {
	case tea.KeyMsg:
		switch msgTyped.String() {
		case "q":
			return m, func() tea.Msg {
				return models.ReturnMsg{}
			}
		}
	case models.RichTask:
		m.task = msgTyped
	case tea.WindowSizeMsg:
		m.Resize(msgTyped.Width, msgTyped.Height)
		// viewport must set content in every sizing
		// example, https://github.com/charmbracelet/bubbletea/blob/master/examples/pager/main.go#L74
		m.Viewport.SetContent(m.genContent())
	}
	m.Viewport, cmd = m.Viewport.Update(msg)

	return m, cmd
}

func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

func (m *Model) Resize(width, height int) {
	if width <= 0 || height <= 0 {
		return
	}
	paneBorder := 1
	const twoSide = 2
	fixedWidth := width - paneBorder*twoSide
	fixedHeight := height - paneBorder*1 // TODO ? 1 or 2
	m.Resizable.Resize(fixedWidth, fixedHeight, styles.PaneStyle.GetBorderStyle())
}
