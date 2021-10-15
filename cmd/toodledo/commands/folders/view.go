package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ViewCmd = &cobra.Command{
	Use:  "view",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := injector.InitApp()
		if err != nil {
			logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
			return
		}

		name := args[0]
		f, err := app.FolderSvc.FindByName(name)
		if err != nil {
			logrus.WithError(err).Fatal()
			return
		}
		fmt.Print(render.Tables4Folder([]*models.Folder{f}))
	},
}
