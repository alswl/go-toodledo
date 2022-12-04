package render

import (
	"bytes"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/olekukonko/tablewriter"
)

// Environments ...
func Environments(cks []*models.EnvironmentWithKey) string {
	buf := new(bytes.Buffer)
	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Key", "Name", "Space", "Project"})
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	for _, x := range cks {
		table.Append([]string{x.Key, x.Name, x.Space, x.Project})
	}
	table.Render()
	return buf.String()
}
