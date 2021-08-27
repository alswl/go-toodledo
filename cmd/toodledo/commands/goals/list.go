package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/goal"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		t := viper.GetString("auth.access_token")
		if t == "" {
			logrus.Error("auth.access_token is empty")
			return
		}
		auth := auth.NewSimpleAuth(t)

		cli := client.NewHTTPClient(strfmt.NewFormats())
		res, err := cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Print(render.Tables4Goal(res.Payload))
	},
}
