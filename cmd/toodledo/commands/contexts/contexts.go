package contexts

import "github.com/spf13/cobra"

var ContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage toodledo contexts",
	Run:   ListCmd.Run,
}

func init() {
	ContextCmd.AddCommand(ListCmd, CreateCmd, DeleteCmd, RenameCmd, ViewCmd)
}
