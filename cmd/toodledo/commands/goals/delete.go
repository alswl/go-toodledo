package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:  "delete",
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
		err = svc.Delete(f.ID)
		if err != nil {
			logrus.WithError(err).Error("failed to delete goal")
			return
		}
		fmt.Println("done")
	},
}
