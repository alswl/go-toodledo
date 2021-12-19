package tasks

import "github.com/spf13/cobra"

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage toodledo tasks",
	Run:   ListCmd.Run,
}

func init() {
	TaskCmd.AddCommand(ListCmd, ViewCmd, CreateCmd, DeleteCmd)
}
