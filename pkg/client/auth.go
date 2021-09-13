package client

import (
	"context"
	"errors"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/go-openapi/runtime"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"os"

	"time"
)

type SimpleAuth struct {
	accessToken string
}

// TODO remove this, using configs instead of accessToken
func NewSimpleAuth(accessToken string) runtime.ClientAuthInfoWriter {
	return &SimpleAuth{accessToken: accessToken}
}

// TODO using configs instead of accessToken
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
	conf, err := ProvideOAuth2Config()
	if err != nil {
		return nil, err
	}

	if token.Expiry.Before(time.Now()) {
		if rt == "" {
			return nil, errors.New("auth.refresh_token is empty")
		}
		newToken, err := regenerate(conf, &token)
		err = SaveTokenToConfig(newToken)
		if err != nil {
			return nil, err
		}
		return NewSimpleAuth(newToken.AccessToken), nil
	}
	return NewSimpleAuth(accessToken), nil
}

// TODO using this
func NewToodledoCli(configs common.Configs) *Toodledo {
	debug := os.Getenv("DEBUG") != "" || os.Getenv("SWAGGER_DEBUG") != ""

	transportConfig := openapiclient.New(DefaultHost, DefaultBasePath, []string{"https"})
	transportConfig.Debug = debug
	return New(transportConfig, strfmt.Default)
}

func (a *SimpleAuth) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetQueryParam("access_token", a.accessToken)
	return nil
}

func ProvideOAuth2Config() (*oauth2.Config, error) {
	clientId := viper.GetString("auth.client_id")
	clientSecret := viper.GetString("auth.client_secret")
	if clientId == "" {
		return nil, errors.New("clientId is required")
	}
	if clientSecret == "" {
		return nil, errors.New("clientSecret is required")
	}
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
	return conf, nil
}

func regenerate(conf *oauth2.Config, oldToken *oauth2.Token) (*oauth2.Token, error) {
	src := conf.TokenSource(context.TODO(), oldToken)
	newToken, err := src.Token()
	if err != nil {
		return nil, errors.New("auth.refresh_token is empty")
	}
	return newToken, nil
}

// TODO move to Configs
func SaveTokenToConfig(tok *oauth2.Token) error {
	// TOTO move to Configs
	viper.Set("auth.access_token", tok.AccessToken)
	viper.Set("auth.expired_at", tok.Expiry.Format(time.RFC3339))
	viper.Set("auth.refresh_token", tok.RefreshToken)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
