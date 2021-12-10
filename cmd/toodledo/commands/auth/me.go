package auth

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Who am i?",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitAccountSvc()
		if err != nil {
			logrus.Fatal(err)
			return
		}

		me, err := svc.Me()
		if err != nil {
			logrus.Error(err)
			return
		}

		// TODO pretty
		out, _ := yaml.Marshal(me)
		fmt.Println((string)(out))
	},
}
