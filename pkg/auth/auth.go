package auth

import (
	"context"
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

func ProvideSimpleAuth() (runtime.ClientAuthInfoWriter, error) {
	// TOTO remove viper dependencies
	accessToken := viper.GetString("auth.access_token")
	rt := viper.GetString("auth.refresh_token")
	if accessToken == "" {
		logrus.Error("auth.access_token is empty")
		return nil, errors.New("auth.access_token is empty")
	}
	expiredAt := viper.GetString("auth.expired_at")
	if expiredAt == "" {
		return nil, errors.New("auth.expired_at is empty")
	}
	at, err := time.Parse(time.RFC3339, expiredAt)
	if err != nil {
		return nil, errors.New("auth.expired_at parse error")
	}

	token := oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: rt,
		Expiry:       at,
	}
	conf := ProvideOAuth2Config()

	if token.Expiry.Before(time.Now()) {
		if rt == "" {
			return nil, errors.New("auth.refresh_token is empty")
		}
		newToken, err := Regenerate(conf, &token)
		err = SaveTokenToConfig(newToken)
		if err != nil {
			return nil, err
		}
		return NewSimpleAuth(newToken.AccessToken), nil
	}
	return NewSimpleAuth(accessToken), nil
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

func Regenerate(conf *oauth2.Config, oldToken *oauth2.Token) (*oauth2.Token, error) {
	src := conf.TokenSource(context.TODO(), oldToken)
	newToken, err := src.Token()
	if err != nil {
		return nil, errors.New("auth.refresh_token is empty")
	}
	return newToken, nil
}

func SaveTokenToConfig(tok *oauth2.Token) error {
	// TOTO remove viper dependencies
	viper.Set("auth.access_token", tok.AccessToken)
	viper.Set("auth.expired_at", tok.Expiry.Format(time.RFC3339))
	viper.Set("auth.refresh_token", tok.RefreshToken)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
