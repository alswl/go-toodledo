package config

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	var Cmd = &cobra.Command{
		Use:   "config",
		Short: "Manage config",
	}
	Cmd.AddCommand(initCmd)
	Cmd.AddCommand(viewCmd)
	return Cmd
}
