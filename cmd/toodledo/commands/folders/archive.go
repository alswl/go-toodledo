package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ArchiveCmd = &cobra.Command{
	Use:  "archive",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		name := args[0]

		f, err := app.FolderSvc.Find(name)
		if err != nil {
			logrus.Error(err)
			return
		}

		newF, err := app.FolderSvc.ArchiveFolder(int(f.ID), false)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Print(render.Tables4Folder([]*models.Folder{newF}))
	},
}
