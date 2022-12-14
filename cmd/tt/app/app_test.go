package app

import (
	"testing"

	"github.com/alswl/go-toodledo/pkg/ui/taskspane"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

func TestInterfaceEqual(t *testing.T) {
	tp := taskspane.Model{}
	var itf tea.Model = &tp
	// assert.True(t, itf == itf)
	assert.Equal(t, itf, itf)
}
