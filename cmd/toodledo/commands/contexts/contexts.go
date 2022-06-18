package contexts

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var ContextCmd = &cobra.Command{
		Use:   "context",
		Short: "Manage toodledo contexts",
	}
	ContextCmd.AddCommand(ListCmd, CreateCmd, DeleteCmd, RenameCmd, ViewCmd)
	return ContextCmd
}
