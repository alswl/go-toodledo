package auth

import (
	"context"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
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

		viper.Set("auth.access_token", tok.AccessToken)
		viper.Set("auth.expired_at", tok.Expiry.Format(time.RFC3339))
		viper.Set("auth.refresh_token", tok.RefreshToken)
		err = viper.WriteConfig()
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println("ok")

	},
}
