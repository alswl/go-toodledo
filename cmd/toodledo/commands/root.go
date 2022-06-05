package commands

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/auth"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/configs"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/contexts"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/folders"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/goals"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/savedsearches"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/tasks"
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

var rootCmd = &cobra.Command{
	Use:              "toodledo",
	TraverseChildren: true,
	Version:          version.Message(),
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("access_token", "", "", "")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toodledo.yaml)")

	_ = viper.BindPFlag("auth.access_token", rootCmd.PersistentFlags().Lookup("access_token"))

	rootCmd.AddCommand(tasks.TaskCmd,
		folders.FolderCmd, contexts.ContextCmd, goals.GoalCmd, savedsearches.SavedSearchCmd,
		auth.AuthCmd, configs.ConfigCmd, completionCmd)
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

		// Search config in home directory with name ".toodledo" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", "toodledo"))
		viper.SetConfigType("yaml")
		viper.SetConfigName(".toodledo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("config file", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
