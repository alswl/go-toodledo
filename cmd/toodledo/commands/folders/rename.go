package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RenameCmd = &cobra.Command{
	Use:  "rename",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitFolderService()
		if err != nil {
			logrus.Fatal(err)
			return
		}
		name := args[0]
		newName := args[1]
		f, err := svc.Rename(name, newName)
		if err != nil {
			logrus.WithError(err).Fatal("rename failed")
			return
		}

		fmt.Println(render.Tables4Folder([]*models.Folder{f}))
	},
}
