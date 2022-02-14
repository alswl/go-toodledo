package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitGoalService()
		if err != nil {
			logrus.WithError(err).Fatal("failed to init goals service")
			return
		}
		all, err := svc.ListAll()
		if err != nil {
			logrus.WithError(err).Fatal("failed to list goals")
			return
		}

		fmt.Println(render.Tables4Goal(all))
	},
}
