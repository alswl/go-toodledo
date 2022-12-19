package sidebar

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch typedMsg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := styles.EmptyStyle.GetFrameSize()
		m.contextList.SetSize(typedMsg.Width-h, typedMsg.Height-v)
	// refresh
	case []models.Context:
		m.Contexts = typedMsg
		for range m.contextList.Items() {
			m.contextList.RemoveItem(0)
		}
		for _, c := range m.Contexts {
			m.contextList.InsertItem(len(m.contextList.Items()), Item{c.ID, c.Name})
		}
	case []models.Folder:
		m.Folders = typedMsg
		for range m.folderList.Items() {
			m.folderList.RemoveItem(0)
		}
		for _, c := range m.Folders {
			m.folderList.InsertItem(len(m.folderList.Items()), Item{c.ID, c.Name})
		}
	case []models.Goal:
		m.Goals = typedMsg
		for range m.goalList.Items() {
			m.goalList.RemoveItem(0)
		}
		for _, c := range m.Goals {
			m.goalList.InsertItem(len(m.goalList.Items()), Item{c.ID, c.Name})
		}
	// change select
	case tea.KeyMsg:
		changed := false
		currentItem0 := m.getVisibleList().SelectedItem()
		currentItem, _ := currentItem0.(Item)
		newItem := currentItem
		switch typedMsg.String() {
		case "h":
			m.updateTab(-1)
			newItem, _ = m.getVisibleList().SelectedItem().(Item)
			changed = true
		case "l":
			m.updateTab(+1)
			newItem, _ = m.getVisibleList().SelectedItem().(Item)
			changed = true
		default:
			// dirty event handle without differ
			cmd = m.updateVisibleList(typedMsg)
			newItem, _ = m.getVisibleList().SelectedItem().(Item)
			if newItem.id != currentItem.id {
				changed = true
			}
		}
		if changed {
			for _, sub := range m.onItemChangeSubscribers {
				err := sub(defaultTabs[m.currentTabIndex], newItem)
				if err != nil {
					m.log.WithError(err).Error("failed to change item")
				}
			}
		}
	}

	return m, cmd
}
