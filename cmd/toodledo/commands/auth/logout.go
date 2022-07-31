package auth

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
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
			err := client.CleanAuthWithViper()
			if err != nil {
				logrus.Error(err)
				return
			}
		},
	}

}
