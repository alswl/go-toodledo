package taskstablepane

import (
	"time"

	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
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
		table.NewColumn(columnKeyCompleted, "[ ]", 3).
			WithStyle(lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#88f"))),
		table.NewFlexColumn(columnKeyTitle, "Title", 0).WithFiltered(true),
		table.NewColumn(columnKeyContext, "Context", 10),
		table.NewColumn(columnKeyPriority, "Priority", 10),
		table.NewColumn(columnKeyStatus, "Status", 10),
		table.NewColumn(columnKeyGoal, "Goal", 10),
		table.NewColumn(columnKeyDue, "DueString", 10),
		table.NewColumn(columnKeyRepeat, "Repeat", 10),
		table.NewColumn(columnKeyLength, "Length", 10),
		table.NewColumn(columnKeyTimer, "Timer", 10),
		table.NewColumn(columnKeyTag, "Tag", 10),
	}
)

func RenderTasksRows(tasks []*models.RichTask) []table.Row {
	var rows []table.Row
	for _, t := range tasks {
		context := t.TheContext
		if context == nil {
			context = &models.Context{Name: "<->"}
		}
		goal := t.TheGoal
		if goal == nil {
			goal = &models.Goal{Name: "<->"}
		}

		title := t.Title
		titleStyled := table.NewStyledCell(title, styles.NoStyle)
		if t.Timeron != 0 {
			titleStyled = table.NewStyledCell(title, styles.ProcessingStyle)
		}

		priority := tpriority.Value2Type(t.Priority)
		priorityStyled := table.NewStyledCell(priority.String(), styles.NoStyle)
		if priority == tpriority.Top {
			priorityStyled = table.NewStyledCell(priority.String(), styles.ErrorStyle)
		} else if priority == tpriority.High {
			priorityStyled = table.NewStyledCell(priority.String(), styles.WarningStyle)
		}

		status := tstatus.Value2Type(t.Status)
		statusStyled := table.NewStyledCell(status.String(), styles.NoStyle)
		if status == tstatus.NextAction {
			statusStyled = table.NewStyledCell(status.String(), styles.ErrorStyle)
		} else if status == tstatus.Active {
			statusStyled = table.NewStyledCell(status.String(), styles.WarningStyle)
		}

		due := t.DueString()
		dueStyled := table.NewStyledCell(due, styles.NoStyle)
		if t.Duedate < time.Now().Unix() {
			dueStyled = table.NewStyledCell(due, styles.ErrorStyle)
		} else if t.Duedate < time.Now().AddDate(0, 0, 1).Unix() {
			dueStyled = table.NewStyledCell(due, styles.WarningStyle)
		}

		rows = append(rows, table.NewRow(
			table.RowData{
				columnKeyID:        t.ID,
				columnKeyCompleted: t.CompletedString(),
				columnKeyTitle:     titleStyled,
				columnKeyContext:   context.Name,
				columnKeyPriority:  priorityStyled,
				columnKeyStatus:    statusStyled,
				columnKeyGoal:      goal.Name,
				columnKeyDue:       dueStyled,
				columnKeyRepeat:    t.RepeatString(),
				columnKeyLength:    t.LengthString(),
				columnKeyTimer:     t.TimerString(),
				columnKeyTag:       t.TagString(),
			},
		))
	}
	return rows
}

func (m *Model) Resize(width, height int) {
	if width <= 0 || height <= 0 {
		return
	}
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
	m.Viewport.SetContent(m.View())
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
