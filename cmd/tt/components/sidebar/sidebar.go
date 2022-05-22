package sidebar

import (
	"github.com/alswl/go-toodledo/cmd/tt/components"
	"github.com/alswl/go-toodledo/cmd/tt/components/common"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO move to styles
var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Item struct {
	id    int64
	title string
}

func (i Item) ID() int64           { return i.id }
func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return "" }
func (i Item) FilterValue() string { return i.title }

var defaultTabs = []string{
	"Contexts",
	"Folders",
	//"Goals",
	//"Priority",
	//"Tags",
	//"Search",
}

type Properties struct {
}

type Model struct {
	components.Focusable
	components.Resizable

	// props
	properties Properties

	// states
	isCollapsed     bool
	currentTabIndex int
	currentTab      string
	Contexts        []models.Context

	// view
	// list has states(selected)
	contextList list.Model
	folderList  list.Model

	// handler
	onChange func(tab string, item Item) error
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) UpdateX(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.contextList.SetSize(msg.Width-h, msg.Height-v)
	case []models.Context:
		m.Contexts = msg
		for i, _ := range m.contextList.Items() {
			m.contextList.RemoveItem(i)
		}
		for _, c := range m.Contexts {
			m.contextList.InsertItem(-1, Item{
				id:    c.ID,
				title: c.Name,
			})
		}
	case tea.KeyMsg:
		changed := false
		var newItem Item
		switch msg.String() {
		case "h":
			m.updateTab(-1)
			// FIXME list changed
			changed = true
		case "l":
			m.updateTab(+1)
			// FIXME list changed
			changed = true
		default:
			// dirty event handle without differ
			list := m.getVisibleList()
			oldItem := list.SelectedItem()
			cmd = m.updateVisibleList(msg)
			newItem0 := list.SelectedItem()
			if newItem0 != nil {
				newItem = newItem0.(Item)
			}
			if oldItem != nil && newItem0 != nil && newItem.id != oldItem.(Item).id {
				changed = true
			}
		}
		if changed {
			m.onChange(defaultTabs[m.currentTabIndex], newItem)
		}
	}

	return m, cmd
}

func (m *Model) updateTab(step int) {
	if step == 0 {
		return
	}

	newIndex := (m.currentTabIndex + step + len(defaultTabs)) % len(defaultTabs)
	m.currentTabIndex = newIndex
}

func (m *Model) getVisibleList() *list.Model {
	tab := defaultTabs[m.currentTabIndex]
	var list *list.Model

	switch tab {
	case "Contexts":
		list = &m.contextList
	case "Folders":
		list = &m.folderList
	default:
		panic("unknown tab")
	}
	return list
}

func (m *Model) updateVisibleList(msg tea.Msg) tea.Cmd {
	tab := defaultTabs[m.currentTabIndex]
	var cmd tea.Cmd

	switch tab {
	case "Contexts":
		m.contextList, cmd = m.contextList.Update(msg)
	case "Folders":
		m.folderList, cmd = m.folderList.Update(msg)
	}
	return cmd
}

func (m Model) View() string {
	// TODO move to styles
	tab := defaultTabs[m.currentTabIndex]
	tabRender := lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		Background(lipgloss.Color("#f0f0f0")).Render("<" + tab + ">")

	m.Viewport.SetContent(
		lipgloss.JoinVertical(
			lipgloss.Left,
			tabRender,
			docStyle.Render(m.getVisibleList().View()),
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

func (m *Model) Resize(width, height int) {
	m.Resizable.Resize(width, height)
}

func InitModel(p Properties,
	onChange func(tab string, item Item) error,
) Model {

	m := Model{
		properties:      p,
		isCollapsed:     false,
		currentTabIndex: 0,
		onChange:        onChange,
		contextList:     common.NewSimpleList(),
		folderList:      common.NewSimpleList(),
	}
	//if len(m.list.Items()) > 0 {
	//	m.list.Select(0)
	//}
	m.Blur()
	return m
}
