package render

import (
	"bytes"
	"github.com/alswl/go-toodledo/pkg/models"

	"github.com/jedib0t/go-pretty/v6/table"
)

func TablesRender(folders []*models.Folder) string {
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
