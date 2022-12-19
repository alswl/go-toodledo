package detail

import (
	"strconv"

	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/ui"

	"github.com/alswl/go-toodledo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// nolint
	switch msgTyped := msg.(type) {
	case tea.KeyMsg:
		switch msgTyped.String() {
		case "q":
			return m, func() tea.Msg {
				return models.ReturnMsg{}
			}
		}
	case models.RichTask:
		m.task = msgTyped
	case tea.WindowSizeMsg:
		m.Resize(msgTyped.Width, msgTyped.Height)
		// viewport must set content in every sizing
		// example, https://github.com/charmbracelet/bubbletea/blob/master/examples/pager/main.go#L74
		m.Viewport.SetContent(m.content())
	}
	m.Viewport, cmd = m.Viewport.Update(msg)

	return m, cmd
}

func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

func (m Model) content() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Link: ")+m.task.Link(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Completed: ")+m.task.CompletedString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("ID: ")+strconv.Itoa(int(m.task.ID)),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Title: ")+m.task.Title,
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Context: ")+m.task.ThatContext().Name,
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Folder: ")+m.task.ThatFolder().Name,
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Goal: ")+m.task.ThatGoal().Name,
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Status: ")+m.task.StatusString(),

		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Due: ")+m.task.DueString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Repeat: ")+m.task.RepeatString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Priority: ")+m.task.PriorityString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Repeat: ")+m.task.RepeatString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Length: ")+m.task.LengthString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Timer: ")+m.task.TimerString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Tag: ")+m.task.TagString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Star: ")+m.task.StarString(),

		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Note: ")+m.task.Note,
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Added: ")+m.task.AddedString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Completed: ")+m.task.CompletedString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Modified: ")+m.task.ModifiedString(),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Via: ")+m.task.Via,
	)
}

func (m Model) View() string {
	m.Viewport.SetContent(m.content())
	style := styles.PaneStyle.Copy()

	return style.Render(m.Viewport.View())
}

func (m *Model) Resize(width, height int) {
	if width <= 0 || height <= 0 {
		return
	}
	paneBorder := 1
	const twoSide = 2
	fixedWidth := width - paneBorder*twoSide
	fixedHeight := height - paneBorder*1 // TODO ? 1 or 2
	m.Resizable.Resize(fixedWidth, fixedHeight, styles.PaneStyle.GetBorderStyle())
}
