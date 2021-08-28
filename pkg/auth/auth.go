package auth

import (
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"time"
)

type SimpleAuth struct {
	accessToken string
}

func NewSimpleAuth(accessToken string) runtime.ClientAuthInfoWriter {
	return &SimpleAuth{accessToken: accessToken}
}

func ProvideSimpleAuth(accessToken string) (runtime.ClientAuthInfoWriter, error) {
	return NewSimpleAuth(accessToken), nil
}

func ProvideAccessToken() (string, error) {
	//conf := ProvideOAuth2Config()

	t := viper.GetString("auth.access_token")
	refresh := viper.GetString("auth.refresh_token")
	if t == "" {
		logrus.Error("auth.access_token is empty")
		return "", errors.New("auth.access_token is empty")
	}
	expiredAt := viper.GetString("auth.expired_at")
	if expiredAt == "" {
		return "", errors.New("auth.expired_at is empty")
	}
	at, err := time.Parse(time.RFC3339, expiredAt)
	if err != nil {
		return "", errors.New("auth.expired_at parse error")
	}
	if time.Now().After(at) {
		if refresh == "" {
			return "", errors.New("auth.refresh_token is empty")
		}

		//ctx := context.Background()
		//client := conf.Client(ctx, &oauth2.Token{
		//	AccessToken:  t,
		//	RefreshToken: refresh,
		//	Expiry:       at,
		//})
		//get, err := client.Get("https://api.toodledo.com/3/account/get.php")
		//client.
		// FIXME @jingchao get token using refresh

		// re auth
		return "", errors.New("auth is expired")
	}
	return t, nil
}

func (a *SimpleAuth) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetQueryParam("access_token", a.accessToken)
	return nil
}

func ProvideOAuth2Config() *oauth2.Config {
	clientId := viper.GetString("auth.client_id")
	clientSecret := viper.GetString("auth.client_secret")
	//scope := "basic%20tasks%20write"
	scopes := []string{"basic", "tasks", "write"}
	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.toodledo.com/3/account/authorize.php",
			TokenURL: "https://api.toodledo.com/3/account/token.php",
		},
	}
	return conf
}
