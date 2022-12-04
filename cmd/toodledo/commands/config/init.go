package config

import (
	"bytes"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type initOptions struct {
	Endpoint     string
	ClientID     string
	ClientSecret string
}

var initOpts initOptions

func NewInitCommand(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Args:  cobra.ExactArgs(2),
		Short: "Initialize config",
		Example: heredoc.Doc(`
			$ toodledo config init <client-id> <client-secret>
`),
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
			config.Auth.ClientID = clientID
			config.Auth.ClientSecret = clientSecret
			bs, _ := yaml.Marshal(config)
			_ = viper.ReadConfig(bytes.NewBuffer(bs))
			err := viper.SafeWriteConfig()
			if err != nil {
				logrus.WithError(err).Fatal("write config file")
				return
			}
			_, _ = fmt.Fprintln(f.IOStreams.Out, "ok")
		},
	}
	cmd.Flags().StringVarP(&initOpts.Endpoint, "endpoint", "e", "https://api.toodledo.com", "toodledo api hostname")
	return cmd
}
