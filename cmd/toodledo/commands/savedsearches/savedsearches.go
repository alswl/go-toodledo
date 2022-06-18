package savedsearches

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var SavedSearchCmd = &cobra.Command{
		Use:     "saved-search",
		Aliases: []string{"ss"},
		Short:   "Manage toodledo saved search",
		// TODO search task by saved-searches
	}
	SavedSearchCmd.AddCommand(ListCmd)
	return SavedSearchCmd
}
