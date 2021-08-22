package commands

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands/folders"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
)

var rootCmd = &cobra.Command{
	Use:              "toodledo",
	TraverseChildren: true,
}

var foldersCmd = &cobra.Command{
	Use: "folder",
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("access_token", "", "", "")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	
	foldersCmd.AddCommand(folders.GetCmd, folders.CreateCmd, folders.RenameCmd, folders.ArchiveCmd,
		folders.UnArchiveCmd)
	
	rootCmd.AddCommand(foldersCmd)

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
