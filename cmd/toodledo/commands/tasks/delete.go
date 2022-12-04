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

func NewDeleteCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Args:  cobra.ExactArgs(1),
		Short: "Delete a task",
		Example: heredoc.Doc(`
			$ toodledo tasks delete 8848
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.TaskSvc

			id, _ := strconv.Atoi(args[0])
			// TODO delete by name
			err = svc.Delete((int64)(id))
			if err != nil {
				logrus.Fatal(err)
				return
			}
			_, _ = fmt.Fprintln(f.IOStreams.Out, "done")
		},
	}
}
