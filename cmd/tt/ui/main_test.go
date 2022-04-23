package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceEqual(t *testing.T) {
	tp := TasksPane{}
	var itf tea.Model = &tp
	assert.True(t, itf == itf)
}
