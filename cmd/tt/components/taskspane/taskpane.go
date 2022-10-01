package taskspane

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/tt/components"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	//columnKeyID       = "id"
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
)

const DefaultTableWidth = 120
const defaultPageSize = 20

var DefaultColumns = []table.Column{
	//table.NewColumn(columnKeyID, "ID", 3).WithFiltered(true).WithStyle(lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#88f"))),
	table.NewColumn(columnKeyCompleted, "[X]", 3).WithFiltered(true).WithStyle(lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#88f"))),
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

func TasksRenderRows(tasks []*models.RichTask) []table.Row {
	var rows []table.Row
	for _, t := range tasks {
		rows = append(rows, table.NewRow(
			table.RowData{
				//columnKeyID:       strconv.Itoa(int(t.ID)),
				columnKeyCompleted: t.Completed(),
				columnKeyTitle:     t.Title,
				columnKeyContext:   t.TheContext.Name,
				columnKeyPriority:  tpriority.PriorityValue2Type(t.Priority),
				columnKeyStatus:    tstatus.StatusValue2Type(t.Status),
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

type Model struct {
	components.Focusable
	components.Resizable

	//app *app.Model

	// TODO delete? all states is in app. sub models is ephemeral, Or maybe props is using here?
	//choices  []string         // items on the to-do list
	//cursor   int              // which to-do list item our cursor is pointing at
	//selected map[int]struct{} // which to-do items are selected

	// properties
	// TODO

	// view
	// FIXME table should be only view mode (without filter mode)
	tableModel table.Model
	tableWidth int
	//props      app.States
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) View() string {
	m.Viewport.SetContent(
		m.tableModel.View(),
	)

	style := styles.PaneStyle.Copy()
	if m.IsFocused() {
		style = styles.FocusedPaneStyle.Copy()
	}
	return style.Render(m.Viewport.View())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// children first, bubble blow up model
	m.tableModel, cmd = m.tableModel.Update(msg)
	// FIXME if table acting on event, then we need get the result, and ignore some msg
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}

	case []*models.RichTask:
		m.tableModel = m.tableModel.WithRows(TasksRenderRows(msg))
	case tea.WindowSizeMsg:
		//top, right, bottom, left := docStyle.GetMargin()
		m.Resize(msg.Width, msg.Height)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) Resize(width, height int) {
	// remove status bar height
	m.Resizable.Resize(width, height, styles.PaneStyle.GetBorderStyle())
	paneBorder := 1

	// remove pane border, table header, and table footer
	// normal style
	//tableHeaderHeight := 3
	//tableFooterHeight := 1
	//tableBorder := 1

	// minimal
	tableHeaderHeight := 1
	tableFooterHeight := 0
	tableBorder := 0
	// patch for bubble-tables bug, table columns calculated
	fixWidth := len(DefaultColumns) - 1

	m.tableModel = m.tableModel.WithPageSize(height - tableBorder*2 - tableHeaderHeight - tableFooterHeight - paneBorder*2).
		WithTargetWidth(width + fixWidth)
}

func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

func (m Model) updateFooter() Model {
	//highlightedRow := m.tableModel.HighlightedRow()

	footerText := fmt.Sprintf(
		"Pg. %d/%d - Currently looking at ID: %s",
		m.tableModel.CurrentPage(),
		m.tableModel.MaxPages(),
		//highlightedRow.Data[columnKeyID],
	)

	m.tableModel = m.tableModel.WithStaticFooter(footerText)
	return m
}

func (m *Model) tableSizeSmall() {
	m.tableWidth = m.tableWidth - 10
	m.tableModel = m.tableModel.WithTargetWidth(m.tableWidth)
}

func (m *Model) tableSizeGreater() {
	m.tableWidth = m.tableWidth + 10
	m.tableModel = m.tableModel.WithTargetWidth(m.tableWidth)
}

func (m *Model) Filter(input textinput.Model) {
	m.tableModel = m.tableModel.WithFilterInput(input)
}

func InitModel(tasks []*models.RichTask) Model {
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
		//WithStaticFooter("").
		//Border(customBorder).
		// TODO flex height
		//WithNoPagination().
		WithPageSize(defaultPageSize).
		//WithNoPagination().
		WithFooterVisibility(false).
		WithTargetWidth(DefaultTableWidth).
		WithKeyMap(keys)

	m := Model{
		//choices:    nil,
		//cursor:     0,
		//selected:   nil,
		tableModel: tb,
		tableWidth: DefaultTableWidth,
		//props:      app.GetStates(),
	}

	m.Blur()

	return m
}
