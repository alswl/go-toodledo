package auth

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/client"
	utilsos "github.com/alswl/go-toodledo/pkg/utils/os"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:  "login",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := injector.InitApp()
		if err == nil {
			me, err := app.AccountSvc.Me()
			if err == nil {
				fmt.Printf("You are already logged in as %s\n", me.Email)
				return
			}
		}

		conf, err := client.NewOAuth2ConfigFromViper()
		if err != nil {
			// XXX test
			logrus.WithError(err).Fatal("init toodledo config, using `toodledo config init` to set up.")
			return
		}

		url := conf.AuthCodeURL("state")
		fmt.Printf("login in your browser in %s\n", url)
		fmt.Println("login in your browser, then copy the param(code) in url to clipboard and run `toodledo auth token YOUR-URL-AFTER-LOGIN`")
		_ = utilsos.OpenInBrowser(url)
	},
}
