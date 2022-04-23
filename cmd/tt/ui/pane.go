package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
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

func (f *Focusable) Blur() {
	f.isFocused = false
}

type ResizeInterface interface {
	Resize(width, height int)
}

type Resizable struct {
	viewport viewport.Model
}

func (r *Resizable) Resize(width, height int) {
	border := lipgloss.NormalBorder()

	r.viewport.Width = width - lipgloss.Width(border.Right+border.Top)
	r.viewport.Height = height - lipgloss.Width(border.Bottom+border.Top)
}
