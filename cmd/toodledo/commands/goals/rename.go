package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RenameCmd = &cobra.Command{
	Use:  "rename",
	Args: cobra.ExactArgs(2),
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
		newName := args[1]
		if name == newName {
			logrus.Error("not changed")
			return
		}

		g, err := svc.FindGoalByName(name)
		if err != nil {
			logrus.Error(err)
			return
		}
		renamed, err := svc.Rename(g.ID, newName)
		fmt.Print(render.Tables4Goal([]*models.Goal{renamed}))
	},
}
