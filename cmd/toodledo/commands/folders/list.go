package folders

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
	// noCache bool
}

// var listOpts = &ListOpts{}

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List folders",
		Example: heredoc.Doc(`
			$ toodledo folder list
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			var svc = app.FolderSvc

			all, err := svc.ListAll()
			if err != nil {
				logrus.WithError(err).Fatal()
				return
			}
			_, _ = fmt.Fprintln(f.IOStreams.Out, render.Tables4Folder(all))
		},
	}
	// cmd.Flags().BoolVarP(&listOpts.noCache, "no-cache", "", false, "do not using cache")
	return cmd
}
