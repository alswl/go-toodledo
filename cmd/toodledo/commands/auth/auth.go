package auth

import (
	"github.com/alswl/go-toodledo/pkg/utils/cmds"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmds.Factory) *cobra.Command {
	var Cmd = &cobra.Command{
		Use:   "auth <command>",
		Short: "Manage authentication",
	}
	Cmd.AddCommand(loginCmd, tokenCmd, statusCmd)
	return Cmd
}
