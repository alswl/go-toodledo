package goals

import "github.com/spf13/cobra"

var GoalCmd = &cobra.Command{
	Use:   "goal",
	Short: "Manage toodledo goals",
	Run:   ListCmd.Run,
}

func init() {
	GoalCmd.AddCommand(ListCmd, CreateCmd, DeleteCmd,
		RenameCmd, ArchiveCmd, ActivateCmd)
}
