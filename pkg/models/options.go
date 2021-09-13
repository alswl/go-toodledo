package models

import "time"

type ToodledoConfigAuth struct {
	AccessToken  string    `mapstructure:"access_token"`
	ClientId     string    `mapstructure:"client_id"`
	ClientSecret string    `mapstructure:"client_secret"`
	ExpiredAt    time.Time `mapstructure:"expired_at"`
	RefreshToken string    `mapstructure:"refresh_token"`
}

type ToodledoConfig struct {
	Auth ToodledoConfigAuth `mapstructure:"auth"`
}
