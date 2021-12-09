package models

// Environment ...
type Environment struct {
	Name            string `json:"name,omitempty"`
	Space           string `json:"space,omitempty"`
	Project         string `json:"project,omitempty"`
	DefaultAssignee string `json:"default-assignee,omitempty"`
	DefaultAssigner string `json:"default-assigner,omitempty"`
}

// EnvironmentWithKey ...
type EnvironmentWithKey struct {
	*Environment
	Key string
}
