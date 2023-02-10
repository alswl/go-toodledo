//go:build integration || tags
// +build integration tags

package viewport

import (
	"testing"

	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

type Model struct {
	ui.Resizable
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	m.Viewport.SetContent("abc")
	style := styles.PaneStyle.Copy()
	return style.Render(m.Viewport.View())
}

func TestViewportWidth(t *testing.T) {
	m := Model{}
	style := styles.PaneStyle.Copy()
	m.Resize(100, 10, style.GetBorderStyle())
	view := m.View()

	t.Log(view)

	assert.Equal(t, `┌───┐
│abc│
│   │
│   │
│   │
│   │
│   │
│   │
│   │
└───┘`, view)
}
