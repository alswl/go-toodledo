package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"os"

	"time"
)

// SimpleAuth ...
type SimpleAuth struct {
	accessToken string
}

func NewSimpleAuth(accessToken string) runtime.ClientAuthInfoWriter {
	return &SimpleAuth{accessToken: accessToken}
}

// AuthenticateRequest ...
func (a *SimpleAuth) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetQueryParam("access_token", a.accessToken)
	return nil
}

// NewAuthFromConfig create auth writer from ToodledoConfig
func NewAuthFromConfig(cfg models.ToodledoConfig) (runtime.ClientAuthInfoWriter, error) {
	accessToken := cfg.AccessToken
	rt := cfg.RefreshToken
	if accessToken == "" {
		logrus.Error("auth.access_token is empty")
		return nil, errors.New("auth.access_token is empty")
	}
	expiredAt := cfg.ExpiredAt
	if expiredAt == "" {
		return nil, errors.New("auth.expired_at is empty")
	}
	at, err := time.Parse(time.RFC3339, expiredAt)
	if err != nil {
		return nil, errors.New("auth.expired_at parse error")
	}

	return NewAuth(cfg.ClientId, cfg.ClientSecret, accessToken, rt, at, noOps)
}

// NewAuth create auth writer by access token and refresh token, it will automatically refresh
func NewAuth(clientId, clientSecret, accessToken, refreshToken string, expiredAt time.Time, saveFn func(newToken *oauth2.Token) error) (runtime.ClientAuthInfoWriter, error) {
	token := oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       expiredAt,
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

	// refresh access token by refresh token
	if token.Expiry.Before(time.Now()) {
		if refreshToken == "" {
			return nil, errors.New("auth.refresh_token is empty")
		}
		newToken, err := regenerate(conf, &token)
		if err != nil {
			return nil, err
		}
		err = saveFn(newToken)
		if err != nil {
			return nil, err
		}
		return NewSimpleAuth(newToken.AccessToken), nil
	}
	return NewSimpleAuth(accessToken), nil
}

// NewToodledoCli ...
func NewToodledoCli() *Toodledo {
	debug := os.Getenv("DEBUG") != "" || os.Getenv("SWAGGER_DEBUG") != ""

	transportConfig := openapiclient.New(DefaultHost, DefaultBasePath, []string{"https"})
	transportConfig.Debug = debug
	return New(transportConfig, strfmt.Default)
}

// NewOAuth2ConfigFromConfigs ...
func NewOAuth2ConfigFromConfigs(cfg models.ToodledoConfig) (*oauth2.Config, error) {
	// TODO remove viper
	clientId := cfg.ClientId
	clientSecret := cfg.ClientSecret
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

// ProvideOAuth2ConfigFromViper ...
func ProvideOAuth2ConfigFromViper() (*oauth2.Config, error) {
	// TODO delete, 3 usage left
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
		return nil, fmt.Errorf("failed to get new token: %s", err)
	}
	return newToken, nil
}

// SaveTokenWithViper save new token to yaml
// TODO refactor
func SaveTokenWithViper(tok *oauth2.Token) error {
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

func noOps(newToken *oauth2.Token) error {
	// XXX no ops
	return nil
}
