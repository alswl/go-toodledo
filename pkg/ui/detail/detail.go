package detail

import (
	"github.com/alswl/go-toodledo/pkg/ui"

	"github.com/alswl/go-toodledo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
)

const defaultColumnLabelWidth = 20

type Model struct {
	ui.Focusable
	ui.Resizable

	task models.RichTask
}

func New(task models.RichTask) *Model {
	return &Model{task: task}
}

func (m Model) Init() tea.Cmd {
	return nil
}
