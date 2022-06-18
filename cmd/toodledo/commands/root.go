package commands

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/auth"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/browser"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/config"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/contexts"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/folders"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/goals"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/savedsearches"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/tasks"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/iostreams"
	"github.com/alswl/go-toodledo/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"

	"os"
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
	// XXX changed, test
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config/toodledo/conf.yaml)")

	_ = viper.BindPFlag("auth.access_token", rootCmd.PersistentFlags().Lookup("access_token"))

	// TODO all subcmd using factory
	rootCmd.AddCommand(tasks.TaskCmd,
		folders.NewCmd(f), contexts.NewCmd(f), goals.NewCmd(f), savedsearches.NewCmd(f),
		browser.NewCmd(f),
		auth.NewCmd(f), config.NewCmd(f), completionCmd)

	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}

		// Search config in ~/.config/toodledo/conf.yaml
		viper.AddConfigPath(path.Join(home, ".config", "toodledo"))
		viper.SetConfigName("conf")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("config file", viper.ConfigFileUsed())
	}
}

func Execute() {
	iostreams.UsingSystem()
	cmd := NewRootCmd(cmdutil.NewFactory())
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
