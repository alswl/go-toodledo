package savedsearches

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
		Short: "List saved searches",
		Example: heredoc.Doc(`
			$ toodledo saved-search list
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.SavedSearchSvc
			all, err := svc.ListAll()
			if err != nil {
				logrus.WithError(err).Fatal("list saved searches")
				return
			}

			fmt.Println(render.Tables4SavedSearches(all))
		},
	}
}
