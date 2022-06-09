package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var viewCmd = &cobra.Command{
	Use: "view",
	Run: func(cmd *cobra.Command, args []string) {
		settings := viper.AllSettings()
		text, _ := yaml.Marshal(settings)
		fmt.Println(text)
	},
}

func init() {
	ConfigCmd.AddCommand(viewCmd)
}
