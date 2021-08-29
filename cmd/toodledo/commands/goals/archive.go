package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ArchiveCmd = &cobra.Command{
	Use:  "archive",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		auth, err := auth.ProvideSimpleAuth()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		name := args[0]

		f, err := service.FindGoalByName(auth, name)
		if err != nil {
			logrus.Error(err)
			return
		}

		newF, err := service.ArchiveGoal(auth, int(f.ID), true)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Print(render.Tables4Goal([]*models.Goal{newF}))
	},
}
