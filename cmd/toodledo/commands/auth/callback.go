package auth

import (
	"context"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:  "token",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]
		if code == "" {
			logrus.WithField("args[0]", code).Error("url format error")
			return
		}
		conf := auth.ProvideOAuth2Config()
		tok, err := conf.Exchange(context.Background(), code)
		if err != nil {
			logrus.Error(err)
			return
		}
		err = auth.SaveTokenToConfig(tok)
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Println("ok")

	},
}
