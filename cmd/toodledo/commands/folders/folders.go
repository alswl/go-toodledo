package folders

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "folder",
		Short: "Manage toodledo folders",
	}
	cmd.AddCommand(NewListCmd(f), NewCreateCmd(f), NewDeleteCmd(f),
		NewRenameCmd(f), NewArchiveCmd(f), NewActivateCmd(f), NewViewCmd(f))
	return cmd
}
