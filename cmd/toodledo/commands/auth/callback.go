package auth

import (
	"context"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:  "token",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]
		if code == "" {
			log.WithField("args[0]", code).Error("url format error")
			return
		}
		conf, err := client.ProvideOAuth2ConfigFromViper()
		if err != nil {
			log.Error(err)
			return
		}
		tok, err := conf.Exchange(context.Background(), code)
		if err != nil {
			log.Error(err)
			return
		}
		err = client.SaveTokenWithViper(tok)
		if err != nil {
			log.Error(err)
			return
		}
		fmt.Println("ok")

	},
}
