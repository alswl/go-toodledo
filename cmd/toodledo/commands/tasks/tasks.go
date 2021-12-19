package tasks

import "github.com/spf13/cobra"

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage toodledo tasks",
	Run:   ListCmd.Run,
}

func init() {
	// FIXME
	EditCmd.Flags().StringP("id", "i", "", "task id")

	TaskCmd.AddCommand(ListCmd, ViewCmd, CreateCmd, DeleteCmd, EditCmd)
}
