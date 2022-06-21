package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

func NewDoneCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:     "done",
		Aliases: []string{"complete"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			_, err := injector.InitApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc, err := injector.InitTaskService()
			if err != nil {
				logrus.WithError(err).Fatal("init task service")
				return
			}

			id, _ := strconv.Atoi(args[0])
			newTReturned, err := svc.Complete(int64(id))
			if err != nil {
				logrus.WithField("id", id).WithError(err).Fatal("complete task")
				return
			}
			fmt.Println(render.Tables4Task([]*models.Task{newTReturned}))
		},
	}
}
