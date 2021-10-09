package render

import (
	"bytes"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums/taskpriorities"
	"github.com/alswl/go-toodledo/pkg/models/enums/taskstatus"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Tables4Task(tasks []*models.Task) string {
	var output string
	buf := bytes.NewBufferString(output)
	t := table.NewWriter()
	t.SetOutputMirror(buf)
	// TODO Color
	//t.SetStyle(table.StyleColoredBright)
	t.SetStyle(table.StyleLight)
	t.Style().Options.DrawBorder = false
	t.AppendHeader(table.Row{"#", "[X]", "Title", "Status", "Context", "Priority", "Folder", "Goal",
		"Due", "Repeat", "Length", "Timer"})
	var rows []table.Row
	for _, x := range tasks {
		completed := "[ ]"
		if x.Completed == 1 {
			completed = "[X]"
		}
		context := "Not Set"
		if x.Context != 0 {
			// TODO query
			context = "TODO"
		}
		folder := "Not Set"
		if x.Folder != 0 {
			// TODO query
			folder = "TODO"
		}
		goal := "Not Set"
		if x.Goal != 0 {
			// TODO query
			goal = "TODO"
		}
		rows = append(rows, table.Row{
			x.ID,
			completed,
			x.Title,
			taskstatus.Value2Type(x.Status),
			context,
			taskpriorities.Value2Type(x.Priority),
			folder,
			goal,
			x.Duedate,
			x.Repeat,
			x.Length,
			x.Timer,
		})
	}
	t.AppendRows(rows)
	t.Render()
	return buf.String()
}
