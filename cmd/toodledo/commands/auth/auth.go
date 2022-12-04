package auth

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "auth",
		Args:  cobra.NoArgs,
		Short: "Manage authentication",
	}
	cmd.AddCommand(NewLoginCmd(f), NewTokenCmd(f), NewStatusCmd(f))
	cmd.AddCommand(NewLogoutCmd(f))
	return cmd
}
