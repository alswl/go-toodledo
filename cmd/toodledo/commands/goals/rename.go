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

func NewRenameCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "rename",
		Args:  cobra.ExactArgs(2),
		Short: "Rename goal",
		Example: heredoc.Doc(`
			$ toodledo goal rename landing-moon new-name
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.GoalSvc
			name := args[0]
			newName := args[1]
			if name == newName {
				logrus.Error("not changed")
				return
			}

			renamed, _ := svc.Rename(name, newName)
			fmt.Print(render.Tables4Goal([]*models.Goal{renamed}))
		},
	}
}
