package commands

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/contexts"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/folders"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/goals"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

var (
	// Used for flags.
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:              "toodledo",
	TraverseChildren: true,
}
var folderCmd = &cobra.Command{
	Use: "folder",
}
var contextCmd = &cobra.Command{
	Use: "context",
}
var goalCmd = &cobra.Command{
	Use: "goal",
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("access_token", "", "", "")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	folderCmd.AddCommand(folders.GetCmd, folders.CreateCmd, folders.DeleteCmd,
		folders.RenameCmd, folders.ArchiveCmd, folders.ActivateCmd)
	contextCmd.AddCommand(contexts.GetCmd, contexts.CreateCmd, contexts.DeleteCmd, contexts.RenameCmd)
	goalCmd.AddCommand(goals.GetCmd, goals.CreateCmd, goals.DeleteCmd,
		goals.RenameCmd, goals.ArchiveCmd, goals.ActivateCmd)

	rootCmd.AddCommand(folderCmd, contextCmd, goalCmd)

	viper.BindPFlag("access_token", rootCmd.PersistentFlags().Lookup("access_token"))

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		// Search config in home directory with name ".toodledo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".toodledo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
