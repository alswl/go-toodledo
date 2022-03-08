package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	Refresh  key.Binding
	PageDown key.Binding
	PageUp   key.Binding
	Help     key.Binding
	Quit     key.Binding
}
