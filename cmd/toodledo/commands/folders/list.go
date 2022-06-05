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
		svc, err := injector.InitFolderCachedService()
		if err != nil {
			logrus.WithError(err).Fatal("init folder service")
			return
		}
		syncer, err := injector.InitSyncer()
		if err != nil {
			logrus.WithError(err).Fatal("init syncer failed")
			return
		}
		//syncer.Start(context.Background())
		err = syncer.SyncOnce()
		if err != nil {
			logrus.WithError(err).Fatal("sync failed")
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
