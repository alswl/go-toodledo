package tasks

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

var output string

var viewCmd = &cobra.Command{
	Use:   "view",
	Args:  cobra.ExactArgs(1),
	Short: "View task",
	Example: heredoc.Doc(`
		$ toodledo tasks view 8848	
	`),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := injector.InitCLIApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc := app.TaskSvc
		taskRichSvc := app.TaskRichSvc

		id, _ := strconv.Atoi(args[0])
		task, err := svc.FindById((int64)(id))
		if err != nil {
			logrus.WithError(err).Fatal("find task failed")
			return
		}

		rt, err := taskRichSvc.Rich(task)
		if err != nil {
			logrus.WithError(err).Fatal("rich task failed")
			return
		}

		switch output {
		case "table":
			fmt.Println(render.Tables4RichTasks([]*models.RichTask{rt}))
		case "yaml":
			output, err := render.Yaml4RichTask(rt)
			if err != nil {
				logrus.WithError(err).Fatal("render yaml failed")
				return
			}
			fmt.Println(output)
		default:
			logrus.Fatal("unknown output type")
		}
	},
}

func init() {
	viewCmd.Flags().StringVarP(&output, "output", "o", "table", "table | yaml")

	(&cobra.Command{
		Use:   "task",
		Short: "Manage toodledo tasks",
	}).AddCommand(viewCmd)
}
