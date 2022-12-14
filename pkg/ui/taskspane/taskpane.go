package taskspane

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/ui"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyID        = "id"
	columnKeyCompleted = "completed"
	columnKeyTitle     = "title"
	columnKeyContext   = "context"
	columnKeyStatus    = "status"
	columnKeyPriority  = "priority"
	columnKeyGoal      = "goal"
	columnKeyDue       = "due"
	columnKeyRepeat    = "repeat"
	columnKeyLength    = "length"
	columnKeyTimer     = "timer"
	columnKeyTag       = "tag"
	defaultTableWidth  = 120
	defaultPageSize    = 20
)

var (
	DefaultColumns = []table.Column{
		table.NewColumn(columnKeyCompleted, "[ ]", 3).WithFiltered(true).
			WithStyle(lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#88f"))),
		table.NewFlexColumn(columnKeyTitle, "Title", 0).WithFiltered(true),
		table.NewColumn(columnKeyContext, "Context", 10),
		table.NewColumn(columnKeyPriority, "Priority", 10),
		table.NewColumn(columnKeyStatus, "Status", 10),
		table.NewColumn(columnKeyGoal, "Goal", 10),
		table.NewColumn(columnKeyDue, "DueString", 10),
		table.NewColumn(columnKeyRepeat, "Repeat", 5),
		table.NewColumn(columnKeyLength, "Length", 5),
		table.NewColumn(columnKeyTimer, "Timer", 5),
		table.NewColumn(columnKeyTag, "Tag", 10),
	}
)

func TasksRenderRows(tasks []*models.RichTask) []table.Row {
	var rows []table.Row
	for _, t := range tasks {
		rows = append(rows, table.NewRow(
			table.RowData{
				columnKeyID:        t.ID,
				columnKeyCompleted: t.CompletedString(),
				columnKeyTitle:     t.Title,
				columnKeyContext:   t.TheContext.Name,
				columnKeyPriority:  tpriority.Value2Type(t.Priority),
				columnKeyStatus:    tstatus.Value2Type(t.Status),
				columnKeyGoal:      t.TheGoal.Name,
				columnKeyDue:       t.DueString(),
				columnKeyRepeat:    t.RepeatString(),
				columnKeyLength:    t.LengthString(),
				columnKeyTimer:     t.TimerString(),
				columnKeyTag:       t.TagString(),
			},
		))
	}
	return rows
}

type parent interface {
	ui.Refreshable
	ui.Notifier
}

func (m *Model) Resize(width, height int) {
	// remove status bar height
	m.Resizable.Resize(width, height, styles.PaneStyle.GetBorderStyle())
	paneBorder := 1

	// remove pane border, table header, and table footer
	// normal style
	// tableHeaderHeight := 3
	// tableFooterHeight := 1
	// tableBorder := 1

	// minimal
	tableHeaderHeight := 1
	tableFooterHeight := 0
	tableBorder := 0
	// patch for bubble-tables bug, table columns calculated
	fixWidth := len(DefaultColumns) - 1

	m.tableModel = m.tableModel.
		WithPageSize(height - tableBorder*2 - tableHeaderHeight - tableFooterHeight - paneBorder*2).
		WithTargetWidth(width + fixWidth)
}

// func (m Model) updateFooter() Model {
//	//highlightedRow := m.tableModel.HighlightedRow()
//
//	footerText := fmt.Sprintf(
//		"Pg. %d/%d",
//		m.tableModel.CurrentPage(),
//		m.tableModel.MaxPages(),
//		//highlightedRow.Data[columnKeyID],
//	)
//
//	m.tableModel = m.tableModel.WithStaticFooter(footerText)
//	return m
//}

// func (m *Model) tableSizeSmall() {
//	m.tableWidth = m.tableWidth - 10
//	m.tableModel = m.tableModel.WithTargetWidth(m.tableWidth)
//}
//
// func (m *Model) tableSizeGreater() {
//	m.tableWidth = m.tableWidth + 10
//	m.tableModel = m.tableModel.WithTargetWidth(m.tableWidth)
//}

func (m *Model) Filter(input textinput.Model) {
	m.tableModel = m.tableModel.WithFilterInput(input)
}
