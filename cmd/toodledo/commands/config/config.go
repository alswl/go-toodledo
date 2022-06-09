package config

import "github.com/spf13/cobra"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage config",
}
