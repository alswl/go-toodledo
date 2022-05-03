package sidebar

import (
	"github.com/alswl/go-toodledo/cmd/tt/ui/components"
	"github.com/alswl/go-toodledo/cmd/tt/ui/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type Model struct {
	components.Focusable
	components.Resizable

	isCollapsed bool
	tabs        []string
	currentTab  string
	items       []string
	list        list.Model
	currentItem string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	// TODO move to styles
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

	m.Viewport.SetContent(
		lipgloss.JoinVertical(
			lipgloss.Left,
			tabsRender,
			docStyle.Render(m.list.View()),
		),
	)
	style := styles.UnfocusedPaneStyle
	if m.IsFocused() {
		style = styles.PaneStyle
	}
	//return style.
	//	Width(m.Viewport.Width).
	//	Height(m.Viewport.Height).
	//	Render(wrap.String(
	//		wordwrap.String(m.Viewport.View(), m.Viewport.Width), m.Viewport.Width),
	//	)
	return style.Render(m.Viewport.View())
}

func InitModel() Model {
	l := list.New([]list.Item{
		item{title: "item1"},
		item{title: "item2"},
		item{title: "item3"},
	}, list.NewDefaultDelegate(), 0, 15)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowTitle(false)
	l.SetShowFilter(false)

	m := Model{
		isCollapsed: false,
		tabs:        []string{"tab1", "tab2"},
		currentTab:  "",
		items:       nil,
		list:        l,
		currentItem: "",
	}
	m.Blur()
	return m
}
