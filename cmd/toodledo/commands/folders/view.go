package folders

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
		Short: "View folder",
		Example: heredoc.Doc(`
			$ toodledo folder view reading
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.FolderSvc
			name := args[0]
			f, err := svc.Find(name)
			if err != nil {
				logrus.WithError(err).Fatal()
				return
			}
			fmt.Println(render.Tables4Folder([]*models.Folder{f}))
		},
	}

}
