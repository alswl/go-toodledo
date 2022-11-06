package components

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FocusableInterface interface {
	Focus()
	Blur()
}

type Focusable struct {
	isFocused bool
}

func (f *Focusable) Focus() {
	f.isFocused = true
}

func (f *Focusable) IsFocused() bool {
	return f.isFocused
}

func (f *Focusable) Blur() {
	f.isFocused = false
}

type ResizeInterface interface {
	Resize(width, height int)
}

// Resizable managed the size of the component,
// Viewpoint is inner viewport of the component, which is used to render the content
// Viewport' size is the size of the component minus the border size
type Resizable struct {
	Height   int
	Width    int
	Viewport viewport.Model
}

func NewResizable(width, height int) Resizable {
	return Resizable{Viewport: viewport.Model{Width: width, Height: height}}
}

func (r *Resizable) Resize(width, height int, border lipgloss.Border) {
	r.Width = width
	r.Height = height
	r.Viewport.Width = width - lipgloss.Width(border.Right+border.Top)
	r.Viewport.Height = height - lipgloss.Width(border.Bottom+border.Top)
}

type Refreshable interface {
	Refresh(isHardRefresh bool) tea.Cmd
}

// Notifier is a component that can notify to the parent component
type Notifier interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}
