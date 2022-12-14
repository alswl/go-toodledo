package taskspane

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
	"github.com/thoas/go-funk"
)

type Model struct {
	ui.Focusable
	ui.Resizable

	// app *app.Model

	// TODO delete? all states is in app. sub models is ephemeral, Or maybe props is using here?
	// choices  []string         // items on the to-do list
	// cursor   int              // which to-do list item our cursor is pointing at
	// selected map[int]struct{} // which to-do items are selected

	// properties
	// TODO

	// view
	// TODO table should be only view mode (without filter mode)
	tableModel table.Model
	tableWidth int

	parent parent
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func InitModel(tasks []*models.RichTask, parent parent) Model {
	keys := table.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down")
	keys.RowUp.SetKeys("k", "up")

	tb := table.New(DefaultColumns).
		WithRows(TasksRenderRows(tasks)).
		HeaderStyle(styles.PaneStyle.Copy().Bold(true).BorderStyle(styles.EmptyBorderStyle)).
		Border(styles.EmptyTableBorderStyle).
		SelectableRows(false).
		Focused(true).
		Filtered(true).
		// TODO disable filter in table
		// WithStaticFooter("").
		// Border(customBorder).
		// TODO flex height
		// WithNoPagination().
		WithPageSize(defaultPageSize).
		// WithNoPagination().
		WithFooterVisibility(false).
		WithTargetWidth(defaultTableWidth).
		WithKeyMap(keys)

	m := Model{
		parent: parent,
		//choices:    nil,
		//cursor:     0,
		//selected:   nil,
		tableModel: tb,
		tableWidth: defaultTableWidth,
		//props:      app.GetStates(),
	}

	m.Blur()

	return m
}

func (m Model) Selected() (int64, error) {
	row := m.tableModel.HighlightedRow()
	if funk.IsZero(row) {
		return 0, nil
	}
	id, _ := row.Data[columnKeyID].(int64)
	return id, nil
}
