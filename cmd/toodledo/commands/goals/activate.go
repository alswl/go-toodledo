package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ActivateCmd = &cobra.Command{
	Use:  "activate",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitGoalService()
		if err != nil {
			logrus.WithError(err).Fatal("init goals service")
			return
		}
		name := args[0]

		g, err := svc.Find(name)
		if err != nil {
			logrus.Error(err)
			return
		}

		newG, err := svc.Archive(int(g.ID), false)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println(render.Tables4Goal([]*models.Goal{newG}))
	},
}
