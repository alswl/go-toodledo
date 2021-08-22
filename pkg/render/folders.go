package render

import (
	"bytes"
	"github.com/alswl/go-toodledo/pkg/models"

	"github.com/jedib0t/go-pretty/v6/table"
)

func Tables4Folder(folders []*models.Folder) string {
	var output string
	buf := bytes.NewBufferString(output)
	t := table.NewWriter()
	t.SetOutputMirror(buf)
	// TODO Color
	//t.SetStyle(table.StyleColoredBright)
	t.SetStyle(table.StyleLight)
	t.Style().Options.DrawBorder = false
	t.AppendHeader(table.Row{"#", "Name", "Archived"})
	var rows []table.Row
	for _, folder := range folders {
		rows = append(rows, table.Row{folder.ID, folder.Name, folder.Archived})
	}
	t.AppendRows(rows)
	t.Render()
	return buf.String()
}

func Tables4Context(contexts []*models.Context) string {
	var output string
	buf := bytes.NewBufferString(output)
	t := table.NewWriter()
	t.SetOutputMirror(buf)
	t.SetStyle(table.StyleLight)
	t.Style().Options.DrawBorder = false
	t.AppendHeader(table.Row{"#", "Name"})
	var rows []table.Row
	for _, x := range contexts {
		rows = append(rows, table.Row{x.ID, x.Name})
	}
	t.AppendRows(rows)
	t.Render()
	return buf.String()
}
func Tables4Goal(goals []*models.Goal) string {
	var output string
	buf := bytes.NewBufferString(output)
	t := table.NewWriter()
	t.SetOutputMirror(buf)
	// TODO Color
	//t.SetStyle(table.StyleColoredBright)
	t.SetStyle(table.StyleLight)
	t.Style().Options.DrawBorder = false
	t.AppendHeader(table.Row{"#", "Name", "Level", "Private", "Archived", "Contributes", "Note"})
	var rows []table.Row
	for _, x := range goals {
		rows = append(rows, table.Row{x.ID, x.Name, x.Level, x.Private, x.Archived, x.Contributes, x.Note})
	}
	t.AppendRows(rows)
	t.Render()
	return buf.String()
}
