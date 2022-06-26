package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	noCache bool
}

var listOpts = &ListOpts{}

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			var svc services.FolderService
			if listOpts.noCache {
				svc = app.FolderSvc
			} else {
				svc = app.FolderCachedSvc
				syncer := app.Syncer
				err = syncer.SyncOnce()
				if err != nil {
					logrus.WithError(err).Fatal("sync failed")
					return
				}
			}

			all, err := svc.ListAll()
			if err != nil {
				logrus.WithError(err).Fatal()
				return
			}
			fmt.Println(render.Tables4Folder(all))
		},
	}
	cmd.Flags().BoolVarP(&listOpts.noCache, "no-cache", "", false, "do not using cache")
	return cmd

}
