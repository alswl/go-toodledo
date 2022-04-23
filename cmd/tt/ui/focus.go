package ui

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
