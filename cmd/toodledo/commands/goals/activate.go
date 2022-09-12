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

func NewActivateCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "activate",
		Args:  cobra.ExactArgs(1),
		Short: "Activate goal",
		Example: heredoc.Doc(`
			$ toodledo goal activate landing-moon
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.GoalSvc
			name := args[0]

			g, err := svc.Find(name)
			if err != nil {
				logrus.Error(err)
				return
			}

			newG, err := svc.Archive(int(g.ID), false)
			if err != nil {
				logrus.Error(err)
				return
			}

			fmt.Println(render.Tables4Goal([]*models.Goal{newG}))
		},
	}
}
