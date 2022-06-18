package goals

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var GoalCmd = &cobra.Command{
		Use:   "goal",
		Short: "Manage toodledo goals",
	}
	GoalCmd.AddCommand(ListCmd, CreateCmd, DeleteCmd,
		RenameCmd, ArchiveCmd, ActivateCmd)
	return GoalCmd
}
