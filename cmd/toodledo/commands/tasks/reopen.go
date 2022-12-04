package tasks

import (
	"fmt"
	"strconv"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewReopenCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "reopen",
		Args:  cobra.ExactArgs(1),
		Short: "Reopen a task",
		Example: heredoc.Doc(`
			$ toodledo tasks reopen 8848
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.TaskSvc

			id, _ := strconv.Atoi(args[0])
			newTReturned, err := svc.UnComplete(int64(id))
			if err != nil {
				logrus.WithField("id", id).WithError(err).Fatal("complete task")
				return
			}
			_, _ = fmt.Fprintln(f.IOStreams.Out, render.Tables4Task([]*models.Task{newTReturned}))
		},
	}
}
