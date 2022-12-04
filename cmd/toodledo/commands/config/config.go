package config

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "config",
		Args:  cobra.NoArgs,
		Short: "Manage config",
	}
	cmd.AddCommand(NewInitCommand(f))
	cmd.AddCommand(NewViewCmd(f))
	return cmd
}
