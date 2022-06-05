package savedsearches

import "github.com/spf13/cobra"

var SavedSearchCmd = &cobra.Command{
	Use:     "saved-search",
	Aliases: []string{"ss"},
	Short:   "Manage toodledo saved search",
	// TODO search task by saved-searches
}
