package models

type FetchTasksMsg struct {
	IsHardRefresh bool
}

// RefreshPropertiesMsg refresh context, goals, folders etc.
type RefreshPropertiesMsg struct {
}

// RefreshTasksMsg refresh tasks.
type RefreshTasksMsg struct {
}

// RefreshUIMsg refreshes UI
// actually, did nothing in Update().
type RefreshUIMsg struct {
}

// ReturnMsg is a message for return parent component, just like tea.Quit.
type ReturnMsg struct {
}

// StatusMsg is a message for status.
type StatusMsg struct {
	Mode         string
	ClearMode    bool
	Message      string
	ClearMessage bool
	Info1        string
	ClearInfo1   bool
	Info2        string
	ClearInfo2   bool
}

type SidebarItemChangeMsg struct {
	tab  string
	item Item
}

func (m SidebarItemChangeMsg) Tab() string {
	return m.tab
}

func (m SidebarItemChangeMsg) Item() Item {
	return m.item
}

func NewSidebarItemChangeMsg(tab string, item Item) *SidebarItemChangeMsg {
	return &SidebarItemChangeMsg{tab: tab, item: item}
}
