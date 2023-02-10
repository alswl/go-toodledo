package sidebar

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	tabName := defaultTabs[m.states.CurrentTabIndex]
	tabRender := styles.NoStyle.Render("<" + tabName + ">")

	m.Viewport.SetContent(
		lipgloss.JoinVertical(
			lipgloss.Left,
			tabRender,
			styles.NoStyle.Render(m.currentList().View()),
		),
	)
	style := styles.PaneStyle.Copy()
	if m.IsFocused() {
		style = styles.FocusedPaneStyle.Copy()
	}
	// return style.
	//	Width(m.Viewport.Width).
	//	Height(m.Viewport.Height).
	//	Render(wrap.String(
	//		wordwrap.String(m.Viewport.View(), m.Viewport.Width), m.Viewport.Width),
	//	)
	return style.Width(m.Viewport.Width).Render(m.Viewport.View())
}

func (m *Model) Resize(width, height int) {
	m.Resizable.Resize(width, height, styles.PaneStyle.GetBorderStyle())
	for _, l := range []*list.Model{&m.contextList, &m.folderList, &m.goalList} {
		l.SetWidth(m.Viewport.Width)
		l.SetHeight(m.Viewport.Height)
	}
	m.Viewport.SetContent(m.View())
}

func (m *Model) currentList() *list.Model {
	tab := defaultTabs[m.states.CurrentTabIndex]
	var l *list.Model

	switch tab {
	case constants.Contexts:
		l = &m.contextList
	case constants.Folders:
		l = &m.folderList
	case constants.Goals:
		l = &m.goalList
	default:
		panic("unknown tab")
	}
	return l
}

func (m Model) GetStates() States {
	return *m.states
}
