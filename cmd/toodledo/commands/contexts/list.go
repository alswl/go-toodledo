package contexts

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List contexts",
		Example: heredoc.Doc(`
			$ toodledo context list
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			// FIXME cached
			svc := app.ContextSvc

			all, err := svc.ListAll()
			if err != nil {
				logrus.Error(err)
				return
			}
			fmt.Println(render.Tables4Context(all))
		},
	}

}
