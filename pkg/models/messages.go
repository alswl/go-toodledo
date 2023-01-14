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

// ReturnMsg is a message for return.
type ReturnMsg struct {
}

type StatusMsg struct {
	Message string
}
