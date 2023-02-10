package statusbar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {

}

func TestView(t *testing.T) {
	m := NewDefault()
	view := m.View()
	assert.Equal(t, " \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m ", view)

	m.Info("status")
	m.SetMode("mode")
	m.SetInfo1("info1")
	m.SetInfo2("info2")
	view = m.View()
	assert.Equal(t, " \x1b[;mmode\x1b[0m  \x1b[;mstatus\x1b[0m  \x1b[;minfo1\x1b[0m  \x1b[;minfo2\x1b[0m ", view)
}

func TestLoadingView(t *testing.T) {
	m := NewDefault()
	m.Info("status")
	m.SetMode("mode")
	m.SetInfo1("info1")
	m.SetInfo2("info2")
	m.StartSpinner()
	view := m.View()
	assert.Equal(t, " \x1b[;mmode\x1b[0m  \x1b[;mstatus\x1b[0m ⣾  \x1b[;minfo1\x1b[0m  \x1b[;minfo2\x1b[0m ", view)
}

func TestViewNoColor(t *testing.T) {
	// TODO set env not works
	// os.Setenv("NO_COLOR", "true")
	// os.Setenv("CI", "yes")
	// t.Setenv("CI", "yes")
	m := NewDefault()
	view := m.View()
	assert.Equal(t, " \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m ", view)
}

func TestViewMode(t *testing.T) {
	m := NewDefault()
	m.mode = "Normal"
	view := m.View()
	assert.Equal(t, " \x1b[;mNormal\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m ", view)
}

func TestMessage(t *testing.T) {
	m := NewDefault()
	m.mode = "Normal"
	view := m.View()
	assert.Equal(t, " \x1b[;mNormal\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m  \x1b[;m\x1b[0m ", view)
}
