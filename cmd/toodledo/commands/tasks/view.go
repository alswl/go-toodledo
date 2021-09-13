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

var ViewCmd = &cobra.Command{
	Use:  "view",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}

		id, _ := strconv.Atoi(args[0])
		task, err := app.TaskSvc.FindById((int64)(id))

		fmt.Println(render.Tables4Task([]*models.Task{task}))
	},
}
