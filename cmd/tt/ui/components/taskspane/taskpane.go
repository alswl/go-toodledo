package taskspane

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/tt/ui/components"
	"github.com/alswl/go-toodledo/cmd/tt/ui/styles"
	"github.com/alswl/go-toodledo/pkg/models"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
	"strconv"
)

const (
	columnKeyID      = "id"
	columnKeyTitle   = "title"
	columnKeyContext = "context"
	columnKeyStatus  = "status"
)

var DefaultColumns = []table.Column{
	table.NewColumn(columnKeyID, "ID", 15).WithFiltered(true).WithStyle(lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#88f"))),
	table.NewFlexColumn(columnKeyTitle, "Title", 100).WithFiltered(true),
	//table.NewColumn(columnKeyTitle, "Title", 50).WithFiltered(true),
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

type Model struct {
	components.Focusable
	components.Resizable

	choices    []string         // items on the to-do list
	cursor     int              // which to-do list item our cursor is pointing at
	selected   map[int]struct{} // which to-do items are selected
	tableModel table.Model
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) View() string {
	m.Viewport.SetContent(
		lipgloss.NewStyle().
			Width(m.Viewport.Width).
			Height(m.Viewport.Height).
			PaddingLeft(0).
			Render(m.tableModel.View()),
	)

	style := styles.UnfocusedPaneStyle
	if m.IsFocused() {
		style = styles.PaneStyle
	}
	return style.
		Width(m.Viewport.Width).
		Height(m.Viewport.Height).
		Render(wrap.String(
			wordwrap.String(m.Viewport.View(), m.Viewport.Width), m.Viewport.Width),
		)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Model) updateFooter() Model {
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

func (m *Model) Resize(width, height int) {
	m.Resizable.Resize(width, height)
	m.tableModel.WithTargetWidth(width - 2)
	// FIXME 10 is not accurate
	// FIXME dynamic set page size is not works
	//m.tableModel.WithPageSize(height - 5)
}

func InitModel(tasks []*models.RichTask) Model {

	keys := table.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down")
	keys.RowUp.SetKeys("k", "up")

	m := Model{
		choices:  nil,
		cursor:   0,
		selected: nil,
		tableModel: table.New(DefaultColumns).
			WithRows(TasksRenderRows(tasks)).
			HeaderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)).
			SelectableRows(false).
			Focused(true).
			Filtered(true).
			//Border(customBorder).
			// TODO flex height
			//WithNoPagination().
			// TODO set 20 first init,
			WithPageSize(20).
			WithKeyMap(keys).
			WithStaticFooter("Footer!"),
		//viewport: viewport.Model{Height: 30, Width: 140},
	}

	m = m.updateFooter()
	m.Blur()

	return m
}
