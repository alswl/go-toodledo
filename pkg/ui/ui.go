package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thoas/go-funk"
)

type FocusableInterface interface {
	// TODO add return tea.Cmd
	Focus()
	// TODO add return tea.Cmd
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
// Viewport' size is the size of the component minus the border size.
type Resizable struct {
	Height   int
	Width    int
	Viewport viewport.Model
}

func (r *Resizable) Resize(width, height int, border lipgloss.Border) {
	r.Width = width
	r.Height = height
	r.Viewport.YPosition = lipgloss.Width(border.Top)
	r.Viewport.Width = width - lipgloss.Width(border.Left+border.Right)
	r.Viewport.Height = height - lipgloss.Width(border.Bottom+border.Top)
}

type Refreshable interface {
	FetchTasks(isHardRefresh bool) tea.Cmd
}

// Notifier is a component that can notify to the parent component.
type Notifier interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

// ContainerizedInterface is a component that can contain other components.
type ContainerizedInterface interface {
	Focused() string
	Next() string
	FocusChild(string)
	Children() []string
}

type Containerized struct {
	focused  string
	children []string
}

func NewContainerized(focused string, children []string) *Containerized {
	return &Containerized{focused: focused, children: children}
}

func (c *Containerized) Focused() string {
	return c.focused
}

func (c *Containerized) Next() string {
	if !funk.ContainsString(c.Children(), c.Focused()) {
		return ""
	}

	currentIdx := funk.IndexOf(c.Children(), c.Focused())
	return c.Children()[(currentIdx+1)%len(c.Children())]
}

func (c *Containerized) Children() []string {
	return c.children
}

func (c *Containerized) FocusChild(child string) {
	c.focused = child
}

// VisibleInterface is a component that can be shown or hidden.
type VisibleInterface interface {
	IsVisible() bool
	Show()
	Hide()
}

type Visible struct {
	visible bool
}

func NewVisible(visible bool) Visible {
	return Visible{visible: visible}
}

func (v *Visible) IsVisible() bool {
	return v.visible
}

func (v *Visible) Show() {
	v.visible = true
}

func (v *Visible) Hide() {
	v.visible = false
}

func (v *Visible) Toggle() {
	v.visible = !v.visible
}
