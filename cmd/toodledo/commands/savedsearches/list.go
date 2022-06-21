package savedsearches

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List saved searches",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := injector.InitApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc, err := injector.InitSavedSearchService()
			if err != nil {
				logrus.WithError(err).Fatal("init saved search service")
				return
			}
			all, err := svc.ListAll()
			if err != nil {
				logrus.WithError(err).Fatal("list saved searches")
				return
			}

			fmt.Println(render.Tables4SavedSearches(all))
		},
	}
}
