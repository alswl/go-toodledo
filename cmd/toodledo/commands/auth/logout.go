package auth

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewLogoutCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Args:  cobra.ExactArgs(0),
		Short: "Logout from Toodledo",
		Example: heredoc.Doc(`
			$ toodledo auth logout
`),
		Run: func(cmd *cobra.Command, args []string) {
			opt, err := injector.InitCLIOption()
			if err != nil {
				logrus.Fatal(err)
				return
			}
			err = common.CleanDatabase(opt.Database)
			if err != nil {
				logrus.Fatal(err)
				return
			}

			err = client.CleanAuthWithViper()
			if err != nil {
				logrus.Error(err)
				return
			}

			_, _ = fmt.Fprintln(f.IOStreams.Out, "Logout successfully")
		},
	}
}
