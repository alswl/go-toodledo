package sidebar

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch typedMsg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := styles.NoStyle.GetFrameSize()
		m.contextList.SetSize(typedMsg.Width-h, typedMsg.Height-v)
	// refresh menus
	case []models.Context:
		m.states.Contexts = typedMsg
		for range m.contextList.Items() {
			m.contextList.RemoveItem(0)
		}
		for _, c := range m.states.Contexts {
			m.contextList.InsertItem(len(m.contextList.Items()), models.NewItem(c.ID, c.Name))
		}
	case []models.Folder:
		m.states.Folders = typedMsg
		for range m.folderList.Items() {
			m.folderList.RemoveItem(0)
		}
		for _, c := range m.states.Folders {
			m.folderList.InsertItem(len(m.folderList.Items()), models.NewItem(c.ID, c.Name))
		}
	case []models.Goal:
		m.states.Goals = typedMsg
		for range m.goalList.Items() {
			m.goalList.RemoveItem(0)
		}
		for _, c := range m.states.Goals {
			m.goalList.InsertItem(len(m.goalList.Items()), models.NewItem(c.ID, c.Name))
		}
	case *States:
		// restore current tab and item
		m.states.CurrentTabIndex = typedMsg.CurrentTabIndex
		currentList := m.currentList()
		id := typedMsg.ItemIndexReadonlyMap[currentList.Title]
		for i, item := range currentList.Items() {
			// nolint: errcheck
			typedItem := item.(models.Item)
			if typedItem.ID() == id {
				currentList.Select(i)
				break
			}
		}
		m.states.IsCollapsed = typedMsg.IsCollapsed

	// change select
	case tea.KeyMsg:
		currentList := m.currentList()
		currentItem0 := currentList.SelectedItem()
		currentItem, _ := currentItem0.(models.Item)
		newItem := currentItem
		// changed indicates whether the main ui should refresh query
		changed := false
		switch typedMsg.String() {
		case "h":
			m.updateTab(-1)
			currentList = m.currentList()
			newItem, _ = currentList.SelectedItem().(models.Item)
			m.states.ItemIndexReadonlyMap[currentList.Title] = newItem.ID()
			changed = true
		case "l":
			m.updateTab(+1)
			currentList = m.currentList()
			newItem, _ = currentList.SelectedItem().(models.Item)
			m.states.ItemIndexReadonlyMap[currentList.Title] = newItem.ID()
			changed = true
		default:
			// other event handle
			cmd = m.updateCurrentList(typedMsg)
			newItem, _ = currentList.SelectedItem().(models.Item)
			m.states.ItemIndexReadonlyMap[currentList.Title] = newItem.ID()
			if newItem.ID() != currentItem.ID() {
				changed = true
			}
		}
		if changed {
			cmd = func() tea.Msg {
				return *models.NewSidebarItemChangeMsg(defaultTabs[m.states.CurrentTabIndex], newItem)
			}
		}
	}

	return m, cmd
}

func (m *Model) updateCurrentList(msg tea.Msg) tea.Cmd {
	tab := defaultTabs[m.states.CurrentTabIndex]
	var cmd tea.Cmd

	switch tab {
	case constants.Contexts:
		update, t := m.contextList.Update(msg)
		m.contextList, cmd = update, t
	case constants.Folders:
		update, t := m.folderList.Update(msg)
		m.folderList, cmd = update, t
	case constants.Goals:
		update, t := m.goalList.Update(msg)
		m.goalList, cmd = update, t
	}
	return cmd
}
func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

func (m *Model) updateTab(step int) {
	if step == 0 {
		return
	}

	newIndex := (m.states.CurrentTabIndex + step + len(defaultTabs)) % len(defaultTabs)
	m.states.CurrentTabIndex = newIndex
}
