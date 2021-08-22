package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ArchiveCmd = &cobra.Command{
	Use:  "archive",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := viper.GetString("auth.access_token")
		if t == "" {
			logrus.Error("auth.access_token is empty")
			return
		}
		auth := auth.NewSimpleAuth(t)
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
