package tasks

import (
	"fmt"
	"strconv"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewStartCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:     "start",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{},
		Short:   "Start a task",
		Example: heredoc.Doc(`
			$ toodledo tasks start 8848
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.TaskSvc

			id, _ := strconv.Atoi(args[0])
			err = svc.Start(int64(id))
			if err != nil {
				logrus.WithField("id", id).WithError(err).Fatal("start task")
				return
			}
			_, _ = fmt.Fprintln(f.IOStreams.Out, "started")
		},
	}
}
