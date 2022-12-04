package contexts

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "context",
		Args:  cobra.NoArgs,
		Short: "Manage toodledo contexts",
	}
	cmd.AddCommand(NewListCmd(f), NewCreateCmd(f), NewDeleteCmd(f), NewRenameCmd(f), NewViewCmd(f))
	return cmd
}
