package auth

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var Cmd = &cobra.Command{
		Use:   "auth",
		Args:  cobra.NoArgs,
		Short: "Manage authentication",
	}
	Cmd.AddCommand(NewLoginCmd(f), NewTokenCmd(f), NewStatusCmd(f))
	return Cmd
}
