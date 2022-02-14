package tasks

import (
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage toodledo tasks",
	Args:  cobra.NoArgs,
}

func init() {
	TaskCmd.AddCommand(listCmd, viewCmd, createCmd, deleteCmd, editCmd, completeCmd, uncompleteCmd)
}
