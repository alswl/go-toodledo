package auth

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Args:  cobra.ExactArgs(0),
	Short: "View authentication status",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := injector.InitApp()
		if err != nil {
			logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		me, err := app.AccountSvc.Me()
		if err != nil {
			logrus.WithError(err).Error("get auth info status")
			return
		}

		// TODO pretty
		out, _ := yaml.Marshal(me)
		fmt.Println((string)(out))
	},
}

func init() {
	Cmd.AddCommand(statusCmd)
}
