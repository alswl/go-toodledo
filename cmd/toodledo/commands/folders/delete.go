package folders

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Args:  cobra.ExactArgs(1),
		Short: "Delete folder",
		Example: heredoc.Doc(`
			$ toodledo folder delete reading
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.FolderSvc

			name := args[0]
			err = svc.Delete(name)
			if err != nil {
				logrus.Fatal(err)
				return
			}
			_, _ = fmt.Fprintln(f.IOStreams.Out, "done")
		},
	}
}
