package auth

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func NewStatusCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Args:  cobra.ExactArgs(0),
		Short: "View authentication status",
		Example: heredoc.Doc(`
			$ toodledo auth status
`),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			me, _, err := app.AccountSvc.CachedMe()
			if err != nil {
				logrus.WithError(err).Error("get auth info status")
				return
			}

			// TODO pretty
			out, _ := yaml.Marshal(me)
			_, _ = fmt.Fprintln(f.IOStreams.Out, (string)(out))
		},
	}
}
