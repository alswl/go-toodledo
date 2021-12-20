package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

var EditCmd = &cobra.Command{
	Use:  "edit",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitTaskService()
		if err != nil {
			logrus.WithError(err).Fatal("failed to init task service")
			return
		}

		id, _ := strconv.Atoi(args[0])
		t, err := svc.FindById(int64(id))
		if err != nil {
			logrus.WithError(err).Fatal("failed to find task")
			return
		}
		newT := models.Task{}
		copier.Copy(&newT, t)
		newT.Title = args[1]
		//TODO to fields, with opt()
		newTReturned, err := svc.Edit(int64(id), &newT)
		if err != nil {
			logrus.WithField("id", id).WithError(err).Fatal("failed to edit task")
			return
		}
		fmt.Println(render.Tables4Task([]*models.Task{newTReturned}))
	},
}
