package goals

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	//noCache bool
	WithArchived bool
}

var listOpts = &ListOpts{}

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List goals",
		Example: heredoc.Doc(`
			$ toodledo goal list
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.GoalSvc

			var goals []*models.Goal
			if listOpts.WithArchived {
				goals, err = svc.ListAllWithArchived()
			} else {
				goals, err = svc.ListAll()
			}
			if err != nil {
				logrus.WithError(err).Fatal("list goals")
				return
			}

			fmt.Println(render.Tables4Goal(goals))
		},
	}
	cmd.Flags().BoolVarP(&listOpts.WithArchived, "with-archived", "a", false, "list all goals including archived")
	return cmd
}
