package savedsearches

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"gopkg.in/yaml.v3"
)

func NewViewCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Args:  cobra.ExactArgs(1),
		Short: "View saved searches",
		Example: heredoc.Doc(`
			$ toodledo saved-search list
		`),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			app, err := injector.InitCLIApp()
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
			filtered, _ := funk.Filter(all, func(x *models.SavedSearch) bool {
				return x.Name == name
			}).([]*models.SavedSearch)
			if len(filtered) == 0 {
				logrus.WithField("name", name).Fatal("saved search not found")
				return
			}
			bs, err := yaml.Marshal(filtered[0])
			if err != nil {
				logrus.WithError(err).Fatal("marshal saved search")
				return
			}
			_, _ = fmt.Fprintln(f.IOStreams.Out, (string)(bs))
		},
	}
}
