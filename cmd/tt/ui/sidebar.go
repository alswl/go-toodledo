package ui

import (
	"github.com/alswl/go-toodledo/cmd/tt/ui/styles"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
	"github.com/thoas/go-funk"
)

// TODO move
var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

type SidebarPane struct {
	Focusable

	isCollapsed bool
	tabs        []string
	currentTab  string
	items       []string
	list        list.Model
	currentItem string
	viewport    viewport.Model
}

func (m SidebarPane) Init() tea.Cmd {
	return nil
}

func (m SidebarPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SidebarPane) View() string {
	tabStyle := lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		Background(lipgloss.Color("#f0f0f0"))
	var tabs = funk.Map(m.tabs, func(x string) string {
		//if m.currentTab == tab {
		//}
		return tabStyle.Render(x)
	}).([]string)

	tabsRender := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

	padding := 0
	m.viewport.SetContent(
		lipgloss.NewStyle().
			Width(m.viewport.Width).
			Height(m.viewport.Height).
			PaddingLeft(padding).
			Render(lipgloss.JoinVertical(
				lipgloss.Left,
				tabsRender,
				docStyle.Render(m.list.View()),
			),
			),
	)
	style := styles.UnfocusedPaneStyle
	if m.isFocused {
		style = styles.PaneStyle
	}
	return style.
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(wrap.String(
			wordwrap.String(m.viewport.View(), m.viewport.Width), m.viewport.Width),
		)
}

func InitSidebarPane() SidebarPane {
	l := list.New([]list.Item{
		item{title: "item1"},
		item{title: "item2"},
		item{title: "item3"},
	}, list.NewDefaultDelegate(), 0, 15)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowTitle(false)
	l.SetShowFilter(false)

	m := SidebarPane{
		isCollapsed: false,
		tabs:        []string{"tab1", "tab2"},
		currentTab:  "",
		items:       nil,
		list:        l,
		currentItem: "",
		viewport:    viewport.Model{Width: 30, Height: 20},
	}
	m.isFocused = false
	return m
}
