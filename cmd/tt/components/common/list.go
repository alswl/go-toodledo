package common

import "github.com/charmbracelet/bubbles/list"

const width = 20
const height = 14

// NewSimpleList return simple and minimal list.Model.
func NewSimpleList() list.Model {
	itemDlgt := list.NewDefaultDelegate()
	itemDlgt.ShowDescription = false
	itemDlgt.SetSpacing(0)

	items := list.New([]list.Item{}, itemDlgt, width, height)
	items.SetShowHelp(false)
	items.SetShowPagination(false)
	items.SetShowTitle(false)
	items.SetShowFilter(false)
	items.SetShowStatusBar(false)

	items.DisableQuitKeybindings()
	return items
}
