package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ArchiveCmd = &cobra.Command{
	Use:  "archive",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitGoalsService()
		if err != nil {
			logrus.WithError(err).Fatal("failed to init goals service")
			return
		}
		name := args[0]

		f, err := svc.FindByName(name)
		if err != nil {
			logrus.Error(err)
			return
		}

		newF, err := svc.Archive(int(f.ID), true)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println(render.Tables4Goal([]*models.Goal{newF}))
	},
}
