package contexts

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
		Short: "Rename a context",
		Example: heredoc.Doc(`
			$ toodledo context rename Work Work-new
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.ContextSvc
			name := args[0]
			newName := args[1]
			if name == newName {
				logrus.Error("not changed")
				return
			}

			c, err := svc.Rename(name, newName)
			if err != nil {
				logrus.Error(err)
				return
			}
			fmt.Println(render.Tables4Context([]*models.Context{c}))
		},
	}
}
