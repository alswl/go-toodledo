package browser

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "browser",
	Short: "Browser toodledo in browser",
}

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	return Cmd
}
