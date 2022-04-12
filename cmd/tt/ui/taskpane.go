package ui

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/models"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
	"strconv"
	"strings"
)

const (
	columnKeyID      = "id"
	columnKeyTitle   = "title"
	columnKeyContext = "context"
	columnKeyStatus  = "status"
)

var DefaultColumns = []table.Column{
	table.NewColumn(columnKeyID, "ID", 15).WithFiltered(true).WithStyle(lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#88f"))),
	table.NewColumn(columnKeyTitle, "Title", 50).WithFiltered(true),
	table.NewColumn(columnKeyContext, "Context", 15),
	table.NewColumn(columnKeyStatus, "Status", 15),
}

func TasksRenderRows(tasks []*models.RichTask) []table.Row {
	var rows []table.Row
	for _, t := range tasks {
		rows = append(rows, table.NewRow(
			table.RowData{
				columnKeyID:      strconv.Itoa(int(t.ID)),
				columnKeyTitle:   t.Title,
				columnKeyContext: t.TheContext.Name,
				columnKeyStatus:  tstatus.StatusValue2Type(t.Status),
			},
		))
	}
	return rows
}

func AllTasksMock() ([]*models.RichTask, error) {
	return []*models.RichTask{
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
			TheGoal:    models.Goal{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "def"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
	}, nil
}

type TasksPane struct {
	choices    []string         // items on the to-do list
	cursor     int              // which to-do list item our cursor is pointing at
	selected   map[int]struct{} // which to-do items are selected
	tableModel table.Model
	viewport   viewport.Model
}

func (m TasksPane) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m TasksPane) View() string {
	borderColor := faintBorder
	border := lipgloss.NormalBorder()
	body := strings.Builder{}

	body.WriteString("Press left/right or page up/down to move pages\n")
	body.WriteString("Press space/enter to select a row, q or ctrl+c to quit\n")

	//var selectedIDs []string
	//for _, row := range m.tableModel.SelectedRows() {
	//	// Slightly dangerous type assumption but fine for demo
	//	selectedIDs = append(selectedIDs, row.Data[columnKeyID].(string))
	//}
	//body.WriteString(fmt.Sprintf("SelectedIDs: %s\n", strings.Join(selectedIDs, ", ")))
	body.WriteString(m.tableModel.View())
	body.WriteString("\n")

	padding := 0
	m.viewport.SetContent(
		lipgloss.NewStyle().
			Width(m.viewport.Width).
			Height(m.viewport.Height).
			PaddingLeft(padding).
			Render(body.String()),
	)

	return lipgloss.NewStyle().
		BorderForeground(borderColor).
		Border(border).
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(wrap.String(
			wordwrap.String(m.viewport.View(), m.viewport.Width), m.viewport.Width),
		)
}

func (m TasksPane) Update(msg tea.Msg) (TasksPane, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.tableModel, cmd = m.tableModel.Update(msg)
	cmds = append(cmds, cmd)

	// We control the footer text, so make sure to update it
	m.updateFooter()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
		//case tea.WindowSizeMsg:
		//top, right, bottom, left := docStyle.GetMargin()
		//m.tableModel.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	return m, tea.Batch(cmds...)
}

func (m TasksPane) Resize(width, height int) {
	border := lipgloss.NormalBorder()

	m.viewport.Width = width - lipgloss.Width(border.Right+border.Top)
	m.viewport.Height = height - lipgloss.Width(border.Bottom+border.Top)
}

func (m TasksPane) updateFooter() TasksPane {
	highlightedRow := m.tableModel.HighlightedRow()

	footerText := fmt.Sprintf(
		"Pg. %d/%d - Currently looking at ID: %s",
		m.tableModel.CurrentPage(),
		m.tableModel.MaxPages(),
		highlightedRow.Data[columnKeyID],
	)

	m.tableModel = m.tableModel.WithStaticFooter(footerText)
	return m
}

func InitialTasksPane() TasksPane {
	//ts, err := AllTasks()
	// FIXME
	ts, err := AllTasksMock()

	if err != nil {
		ts = []*models.RichTask{}
	}

	keys := table.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down")
	keys.RowUp.SetKeys("k", "up")

	m := TasksPane{
		choices:  nil,
		cursor:   0,
		selected: nil,
		tableModel: table.New(DefaultColumns).
			WithRows(TasksRenderRows(ts)).
			HeaderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)).
			SelectableRows(false).
			Focused(true).
			Filtered(true).
			//Border(customBorder).
			WithKeyMap(keys).
			WithStaticFooter("Footer!").
			WithPageSize(20),
		viewport: viewport.Model{Height: 30, Width: 140},
	}

	m = m.updateFooter()

	return m
}
