package components

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

func (f *Focusable) IsFocused() bool {
	return f.isFocused
}

func (f *Focusable) Blur() {
	f.isFocused = false
}

type ResizeInterface interface {
	Resize(width, height int)
}

type Resizable struct {
	Viewport viewport.Model
}

func NewResizable(width, height int) Resizable {
	return Resizable{Viewport: viewport.Model{Width: width, Height: height}}
}

func (r *Resizable) Resize(width, height int) {
	border := lipgloss.NormalBorder()

	r.Viewport.Width = width - lipgloss.Width(border.Right+border.Top)
	r.Viewport.Height = height - lipgloss.Width(border.Bottom+border.Top)
}
