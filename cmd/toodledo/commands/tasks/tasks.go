package tasks

import "github.com/spf13/cobra"

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage toodledo tasks",
}

func init() {
	TaskCmd.AddCommand(ListCmd, ViewCmd)
}
