package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitFolderService()
		if err != nil {
			logrus.Fatal(err)
			return
		}

		all, err := svc.ListAll()
		if err != nil {
			logrus.WithError(err).Fatal()
			return
		}
		fmt.Println(render.Tables4Folder(all))
	},
}
