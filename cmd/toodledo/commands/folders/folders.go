package folders

import (
	"github.com/spf13/cobra"
)

var FolderCmd = &cobra.Command{
	Use:   "folder",
	Short: "Manage toodledo folders",
}

func init() {
	FolderCmd.AddCommand(ListCmd, CreateCmd, DeleteCmd,
		RenameCmd, ArchiveCmd, ActivateCmd, ViewCmd)
}
