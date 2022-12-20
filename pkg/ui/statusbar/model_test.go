package statusbar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	m := NewDefault()
	view := m.View()
	assert.Equal(t, " \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m ", view)

	m.SetStatus("status")
	m.SetMode("mode")
	m.SetInfo1("info1")
	m.SetInfo2("info2")
	view = m.View()
	assert.Equal(t, " \x1b[;mmode\x1b[0m  \x1b[;mstatus\x1b[0m  \x1b[;minfo1\x1b[0m  \x1b[;minfo2\x1b[0m ", view)
}

func TestLoadingView(t *testing.T) {
	m := NewDefault()
	m.SetStatus("status")
	m.SetMode("mode")
	m.SetInfo1("info1")
	m.SetInfo2("info2")
	m.StartSpinner()
	view := m.View()
	assert.Equal(t, " \x1b[;mmode\x1b[0m  \x1b[;mstatus\x1b[0m ⣾  \x1b[;minfo1\x1b[0m  \x1b[;minfo2\x1b[0m ", view)
}
