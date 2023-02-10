package taskstablepane

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

	// properties
	// TODO

	// view
	// TODO table should be only view mode (without filter mode)
	tableModel table.Model
	tableWidth int
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func InitModel(tasks []*models.RichTask, width, height int) Model {
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
		//parent: parent,
		//choices:    nil,
		//cursor:     0,
		//selected:   nil,
		tableModel: tb,
		tableWidth: defaultTableWidth,
		//props:      app.GetStates(),
	}
	m.Resize(width, height)

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

func (m Model) CurrentAndTotalPage() (int, int) {
	rows := m.tableModel.TotalRows()
	size := m.tableModel.PageSize()
	all := (rows + size - 1) / size
	if all == 0 {
		all = 1 // first page is 1
	}
	return m.tableModel.CurrentPage(), all
}
