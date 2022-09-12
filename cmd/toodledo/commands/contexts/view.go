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

func NewViewCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Args:  cobra.ExactArgs(1),
		Short: "View config",
		Example: heredoc.Doc(`
			$ toodledo context view Work
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.ContextSvc
			name := args[0]
			f, err := svc.Find(name)
			if err != nil {
				logrus.WithError(err).Fatal()
				return
			}
			fmt.Println(render.Tables4Context([]*models.Context{f}))
		},
	}
}
