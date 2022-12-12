package root

import (
	"fmt"
	"os"

	"github.com/alswl/go-toodledo/pkg/iostreams"

	"github.com/alswl/go-toodledo/cmd/toodledo/commands"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/auth"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/browse"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/config"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/contexts"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/folders"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/goals"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/savedsearches"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/tasks"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
)

func NewRootCmd(f *cmdutil.Factory) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:              "toodledo",
		TraverseChildren: true,
		Version:          version.Message(),
	}

	rootCmd.PersistentFlags().StringP("access_token", "", "", "")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config/toodledo/conf.yaml)")

	_ = viper.BindPFlag(common.AuthAccessToken, rootCmd.PersistentFlags().Lookup("access_token"))

	// TODO add Environment
	rootCmd.AddCommand(tasks.NewCmd(f), folders.NewCmd(f), contexts.NewCmd(f), goals.NewCmd(f),
		savedsearches.NewCmd(f), browse.NewCmd(f),
		auth.NewCmd(f), config.NewCmd(f), commands.NewCompletionCmd(f))

	return rootCmd
}

func init() {
	cobra.OnInitialize(func() {
		common.InitViper(cfgFile, ".config/toodledo", "conf")
	})
}

func Execute() {
	iostreams.UsingSystem()
	cmd := NewRootCmd(cmdutil.NewFactory())
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
