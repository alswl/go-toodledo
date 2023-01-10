package detail

import (
	"fmt"
	"strconv"

	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	m.Viewport.SetContent(m.genContent())
	style := styles.PaneStyle.Copy()

	return style.Render(m.Viewport.View())
}

func (m Model) genContent() string {
	repeatRow := "Repeat:"
	if m.task.Repeat != "" {
		repeatRow = fmt.Sprintf("Repeat: %s (%s)", m.task.RepeatString(), m.task.Repeat)
	}
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
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render(repeatRow),
		lipgloss.NewStyle().Width(defaultColumnLabelWidth).Render("Priority: ")+m.task.PriorityString(),
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
