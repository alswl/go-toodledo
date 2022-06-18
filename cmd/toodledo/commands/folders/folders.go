package folders

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var FolderCmd = &cobra.Command{
		Use:   "folder",
		Short: "Manage toodledo folders",
	}
	FolderCmd.AddCommand(ListCmd, CreateCmd, DeleteCmd,
		RenameCmd, ArchiveCmd, ActivateCmd, ViewCmd)
	return FolderCmd
}
