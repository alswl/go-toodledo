package goals

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	noCache bool
}

var listOpts = &ListOpts{}

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List goals",
		Example: heredoc.Doc(`
			$ toodledo goal list
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.GoalSvc
			if !listOpts.noCache {
				svc = app.GoalCachedSvc
				syncer := app.Syncer
				err = syncer.SyncOnce()
				if err != nil {
					logrus.Fatal("sync failed")
					return
				}
			}
			all, err := svc.ListAll()
			if err != nil {
				logrus.WithError(err).Fatal("list goals")
				return
			}

			fmt.Println(render.Tables4Goal(all))
		},
	}
	cmd.Flags().BoolVarP(&listOpts.noCache, "no-cache", "", false, "do not using cache")
	return cmd
}
