package config

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type initOptions struct {
	Endpoint     string
	ClientID     string
	ClientSecret string
}

var initOpts initOptions

var initCmd = &cobra.Command{
	Use:  "init <client-id> <client-secret>",
	Args: cobra.ExactArgs(2),
	Long: heredoc.Doc(`
init toodledo config,
you can register your own app here: https://api.toodledo.com/3/account/doc_register.php 
`),
	Run: func(cmd *cobra.Command, args []string) {
		clientID := args[0]
		clientSecret := args[1]

		if viper.ConfigFileUsed() != "" {
			logrus.Fatal("config file already exists, please view $HOME/.config/toodledo/conf.yaml")
			return
		}

		config := models.NewToodledoCliConfig()
		if initOpts.Endpoint != "" {
			config.Endpoint = initOpts.Endpoint
		}
		config.Auth.ClientId = clientID
		config.Auth.ClientSecret = clientSecret
		// FIXME not works, "" in top level
		viper.Set("", config)
		err := viper.SafeWriteConfig()
		if err != nil {
			logrus.WithError(err).Fatal("write config file")
			return
		}
		fmt.Println("ok")
	},
}

func init() {
	initCmd.Flags().StringVarP(&initOpts.Endpoint, "endpoint", "e", "https://api.toodledo.com", "toodledo api hostname")

	Cmd.AddCommand(initCmd)
}
