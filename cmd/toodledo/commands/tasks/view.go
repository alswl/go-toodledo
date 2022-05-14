package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

var viewCmd = &cobra.Command{
	Use:  "view",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitTaskService()
		if err != nil {
			logrus.Fatal(err)
			return
		}
		taskRichSvc, err := injector.InitTaskRichService()
		if err != nil {
			logrus.WithError(err).Fatal("init task rich service failed")
			return
		}

		id, _ := strconv.Atoi(args[0])
		task, err := svc.FindById((int64)(id))
		if err != nil {
			logrus.WithError(err).Fatal("find task failed")
			return
		}

		rts, _ := taskRichSvc.RichThem([]*models.Task{task})
		fmt.Println(render.Tables4RichTasks(rts))
	},
}
