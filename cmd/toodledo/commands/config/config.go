package config

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var Cmd = &cobra.Command{
		Use:   "config",
		Args:  cobra.NoArgs,
		Short: "Manage config",
	}
	Cmd.AddCommand(NewInitCommand(f))
	Cmd.AddCommand(NewViewCmd(f))
	return Cmd
}
