package folders

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewRenameCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "rename",
		Args:  cobra.ExactArgs(2),
		Short: "Rename folder",
		Example: heredoc.Doc(`
			$ toodledo folder rename reading new-name
		`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			svc := app.FolderSvc
			name := args[0]
			newName := args[1]
			folder, err := svc.Rename(name, newName)
			if err != nil {
				logrus.WithError(err).Fatal("rename failed")
				return
			}

			_, _ = fmt.Fprintln(f.IOStreams.Out, render.Tables4Folder([]*models.Folder{folder}))
		},
	}
}
