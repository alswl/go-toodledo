package auth

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/client0"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewTokenCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "token",
		Args:  cobra.ExactArgs(1),
		Short: "Get access token from code",
		Example: heredoc.Doc(`
			$ toodledo auth token YOUR-TOKEN-IN-THE-URL
`),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]
			if code == "" {
				log.WithField("args[0]", code).Error("url format error")
				return
			}
			conf, err := client0.NewOAuth2ConfigFromViper()
			if err != nil {
				log.Error(err)
				return
			}
			tok, err := conf.Exchange(context.Background(), code)
			if err != nil {
				log.Error(err)
				return
			}
			err = client0.SaveTokenWithViper(tok)
			if err != nil {
				log.Error(err)
				return
			}
			app, err := injector.InitCLIApp()
			if err != nil {
				log.Error(err)
				return
			}
			me, err := app.AccountSvc.Me()
			if err != nil {
				log.Error(err)
				return
			}
			err = client0.SaveUserIDWithViper(me.Userid)
			if err != nil {
				log.Error(err)
				return
			}
			_, _ = fmt.Fprintf(f.IOStreams.Out, "You are logged in as %s(%s)\n", me.Userid, me.Email)
			_, _ = fmt.Fprintln(f.IOStreams.Out, "ok")
		},
	}
}
