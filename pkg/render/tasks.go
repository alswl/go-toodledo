package render

import (
	"bytes"
	"github.com/alswl/go-toodledo/pkg/models"
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
	t.AppendHeader(table.Row{"#", "X", "Title", "Status", "Context", "Priority", "Folder", "Goal",
		"Due", "Repeat", "Length", "Timer"})
	var rows []table.Row
	for _, x := range tasks {
		rows = append(rows, table.Row{
			x.ID,
			x.Completed,
			x.Title,
			taskstatus.Value2Type(x.Status),
			x.Context,
			x.Priority,
			x.Folder,
			x.Goal,
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
