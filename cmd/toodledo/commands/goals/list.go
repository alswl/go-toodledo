package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/goal"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		auth, err := client.ProvideSimpleAuth()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}

		cli := client.NewHTTPClient(strfmt.NewFormats())
		res, err := cli.Goal.GetGoalsGetPhp(goal.NewGetGoalsGetPhpParams(), auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Print(render.Tables4Goal(res.Payload))
	},
}
