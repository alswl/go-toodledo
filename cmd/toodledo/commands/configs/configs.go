package configs

import "github.com/spf13/cobra"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage config",
}

func init() {
	ConfigCmd.AddCommand(viewCmd)
}
