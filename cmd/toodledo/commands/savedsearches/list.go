package savedsearches

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
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
			logrus.WithError(err).Fatal("failed to init saved search service")
			return
		}
		all, err := svc.ListAll()
		if err != nil {
			logrus.WithError(err).Fatal("failed to list saved searches")
			return
		}

		fmt.Println(render.Tables4SavedSearches(all))
	},
}
