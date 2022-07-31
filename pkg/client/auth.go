package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/models"
	utilsviper "github.com/alswl/go-toodledo/pkg/utils/viper"
	"github.com/go-openapi/runtime"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"os"

	"time"
)

const AuthAccessToken = "auth.access_token"
const AuthExpiredAt = "auth.expired_at"
const AuthRefreshToken = "auth.refresh_token"

// SimpleAuth ...
type SimpleAuth struct {
	accessToken string
}

// NewAuthByToken is simple runtime.ClientAuthInfoWriter with accessToken
func NewAuthByToken(accessToken string) runtime.ClientAuthInfoWriter {
	return &SimpleAuth{accessToken: accessToken}
}

// AuthenticateRequest ...
func (a *SimpleAuth) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	_ = request.SetQueryParam("access_token", a.accessToken)
	return nil
}

// NewAuthFromConfig create auth writer from ToodledoConfig
func NewAuthFromConfig(cfg models.ToodledoConfig) (runtime.ClientAuthInfoWriter, error) {
	accessToken := cfg.AccessToken
	rt := cfg.RefreshToken
	if accessToken == "" {
		logrus.WithField("key", AuthAccessToken).Error("empty")
		return nil, fmt.Errorf("%s is empty", AuthAccessToken)
	}
	expiredAt := cfg.ExpiredAt
	if expiredAt == "" {
		return nil, fmt.Errorf("%s is empty", AuthExpiredAt)
	}
	at, err := time.Parse(time.RFC3339, expiredAt)
	if err != nil {
		return nil, fmt.Errorf("%s is invalid", AuthExpiredAt)
	}

	return NewAuthWithRefresh(cfg.ClientId, cfg.ClientSecret, accessToken, rt, at, SaveTokenWithViper)
}

// NewAuthWithRefresh create auth writer by access token and refresh token, it will automatically refresh
func NewAuthWithRefresh(clientId, clientSecret, accessToken, refreshToken string, expiredAt time.Time,
	saveFn func(newToken *oauth2.Token) error) (runtime.ClientAuthInfoWriter, error) {
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
	if !token.Expiry.Before(time.Now()) {
		logrus.WithField("expired_at", token.Expiry).Debug("not expired, using token from config")
		return NewAuthByToken(token.AccessToken), nil
	}

	if refreshToken == "" {
		return nil, fmt.Errorf("%s is empty", AuthRefreshToken)
	}
	newToken, err := regenerate(conf, &token)
	if err != nil {
		return nil, err
	}
	err = saveFn(newToken)
	if err != nil {
		return nil, err
	}
	return NewAuthByToken(newToken.AccessToken), nil
}

// NewToodledo ...
func NewToodledo() *Toodledo {
	debug := os.Getenv("DEBUG") != "" || os.Getenv("SWAGGER_DEBUG") != ""

	transportConfig := openapiclient.New(DefaultHost, DefaultBasePath, []string{"https"})
	transportConfig.Debug = debug
	return New(transportConfig, strfmt.Default)
}

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

func NewOAuth2ConfigFromViper() (*oauth2.Config, error) {
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
			// FIXME endpoint configurable
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
		logrus.WithField("err", err).Error("regenerate token failed")
		return nil, fmt.Errorf("get new token: %s", err)
	}
	return newToken, nil
}

// SaveTokenWithViper save new token to yaml
// TODO refactor
func SaveTokenWithViper(tok *oauth2.Token) error {
	// TOTO move to Configs
	viper.Set(AuthAccessToken, tok.AccessToken)
	viper.Set(AuthExpiredAt, tok.Expiry.Format(time.RFC3339))
	viper.Set(AuthRefreshToken, tok.RefreshToken)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func CleanAuthWithViper() error {
	err := utilsviper.Unset(AuthAccessToken)
	if err != nil {
		return err
	}
	err = utilsviper.Unset(AuthRefreshToken)
	if err != nil {
		return err
	}
	err = utilsviper.Unset(AuthExpiredAt)
	if err != nil {
		return err
	}

	return nil
}
