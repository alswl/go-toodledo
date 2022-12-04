package tasks

import (
	"fmt"
	"strconv"

	"github.com/alswl/go-toodledo/pkg/cmdutil"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var output string

func NewViewCmd(f *cmdutil.Factory) *cobra.Command {
	command := &cobra.Command{
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
			task, err := svc.FindByID((int64)(id))
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
				_, _ = fmt.Fprintln(f.IOStreams.Out, render.Tables4RichTasks([]*models.RichTask{rt}))
			case "yaml":
				bs, ierr := render.Yaml4RichTask(rt)
				if ierr != nil {
					logrus.WithError(ierr).Fatal("render yaml failed")
					return
				}
				_, _ = fmt.Fprintln(f.IOStreams.Out, bs)
			default:
				logrus.Fatal("unknown output type")
			}
		},
	}
	command.Flags().StringVarP(&output, "output", "o", "table", "table | yaml")
	return command
}
