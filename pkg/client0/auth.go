package client0

import (
	"context"
	"errors"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/common"

	utilsviper "github.com/alswl/go-toodledo/pkg/utils/viper"
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

// NewAuthByToken is simple runtime.ClientAuthInfoWriter with accessToken.
// TODO accessToken cannot refresh
func NewAuthByToken(accessToken string) runtime.ClientAuthInfoWriter {
	return &SimpleAuth{accessToken: accessToken}
}

func (a *SimpleAuth) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	_ = request.SetQueryParam("access_token", a.accessToken)
	return nil
}

// NewAuthFromConfig create auth writer from ToodledoConfig.
func NewAuthFromConfig(cfg common.ToodledoConfig) (runtime.ClientAuthInfoWriter, error) {
	accessToken := cfg.AccessToken
	rt := cfg.RefreshToken
	if accessToken == "" {
		logrus.WithField("key", common.AuthAccessToken).Error("empty")
		return nil, fmt.Errorf("%s is empty", common.AuthAccessToken)
	}
	expiredAt := cfg.ExpiredAt
	if expiredAt == "" {
		return nil, fmt.Errorf("%s is empty", common.AuthExpiredAt)
	}
	at, err := time.Parse(time.RFC3339, expiredAt)
	if err != nil {
		return nil, fmt.Errorf("%s is invalid", common.AuthExpiredAt)
	}

	return NewAuthWithRefresh(cfg.ClientID, cfg.ClientSecret, accessToken, rt, at, SaveTokenWithViper)
}

// NewAuthWithRefresh create auth writer by access token and refresh token, it will automatically refresh.
func NewAuthWithRefresh(clientID, clientSecret, accessToken, refreshToken string, expiredAt time.Time,
	saveFn func(newToken *oauth2.Token) error) (runtime.ClientAuthInfoWriter, error) {
	token := oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       expiredAt,
	}
	scopes := []string{"basic", "tasks", "write"}
	conf := &oauth2.Config{
		ClientID:     clientID,
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
		return nil, fmt.Errorf("%s is empty", common.AuthRefreshToken)
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

func NewOAuth2ConfigFromConfigs(cfg common.ToodledoConfig) (*oauth2.Config, error) {
	// TODO remove viper
	clientID := cfg.ClientID
	clientSecret := cfg.ClientSecret
	scopes := []string{"basic", "tasks", "write"}
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.toodledo.com/3/account/authorize.php",
			TokenURL: "https://api.toodledo.com/3/account/token.php",
		},
	}
	return conf, nil
}

// NewOAuth2ConfigFromViper get oauth2 config from viper.
// oauth2.Config presumes a login user.
func NewOAuth2ConfigFromViper() (*oauth2.Config, error) {
	// TODO delete, 3 usage left
	clientID := viper.GetString("auth.client_id")
	clientSecret := viper.GetString("auth.client_secret")
	if clientID == "" {
		return nil, errors.New("clientID is required")
	}
	if clientSecret == "" {
		return nil, errors.New("clientSecret is required")
	}
	scopes := []string{"basic", "tasks", "write"}
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			// TODO endpoint configurable
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
		return nil, fmt.Errorf("get new token: %w", err)
	}
	return newToken, nil
}

// SaveTokenWithViper save new token to yaml
// TODO refactor
func SaveTokenWithViper(tok *oauth2.Token) error {
	// TOTO move to Configs
	viper.Set(common.AuthAccessToken, tok.AccessToken)
	viper.Set(common.AuthExpiredAt, tok.Expiry.Format(time.RFC3339))
	viper.Set(common.AuthRefreshToken, tok.RefreshToken)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

// SaveUserIDWithViper save user id to yaml.
func SaveUserIDWithViper(userID string) error {
	viper.Set(common.AuthUserID, userID)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func CleanAuthWithViper() error {
	_ = utilsviper.Unset(common.AuthAccessToken)
	_ = utilsviper.Unset(common.AuthRefreshToken)
	_ = utilsviper.Unset(common.AuthExpiredAt)
	_ = utilsviper.Unset(common.AuthUserID)
	return nil
}
