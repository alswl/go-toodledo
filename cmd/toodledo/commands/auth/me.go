package auth

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/account"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Who am i?",
	Run: func(cmd *cobra.Command, args []string) {
		auth, err := client.ProvideSimpleAuth()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		cli := client.NewHTTPClient(strfmt.NewFormats())
		p := account.NewGetAccountGetPhpParams()
		resp, err := cli.Account.GetAccountGetPhp(p, auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		out, _ := yaml.Marshal(resp.Payload)
		fmt.Print((string)(out))
	},
}
