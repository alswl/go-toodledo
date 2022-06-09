package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type initOptions struct {
	Hostname     string
	ClientID     string
	ClientSecret string
}

var initOpts initOptions

var initCmd = &cobra.Command{
	Use: "init --client-id <client-id> --client-secret <client-secret>",
	Run: func(cmd *cobra.Command, args []string) {
		// XXX impl init with params

		settings := viper.AllSettings()
		text, _ := yaml.Marshal(settings)
		fmt.Println(text)
	},
}

func init() {
	initCmd.Flags().StringVarP(&initOpts.Hostname, "hostname", "h", "https://api.toodledo.com", "toodledo api hostname")
	initCmd.MarkFlagRequired("hostname")
	initCmd.Flags().StringVarP(&initOpts.ClientID, "client-id", "i", "", "toodledo client id")
	initCmd.MarkFlagRequired("client-id")
	initCmd.Flags().StringVarP(&initOpts.ClientSecret, "client-secret", "s", "", "toodledo client secret")
	initCmd.MarkFlagRequired("client-secret")

	ConfigCmd.AddCommand(initCmd)
}
