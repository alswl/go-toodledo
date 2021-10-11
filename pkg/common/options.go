package common

import (
	"time"
)

type ToodledoConfigAuth struct {
	AccessToken  string    `mapstructure:"access_token"`
	ClientId     string    `mapstructure:"client_id"`
	ClientSecret string    `mapstructure:"client_secret"`
	ExpiredAt    time.Time `mapstructure:"expired_at"`
	RefreshToken string    `mapstructure:"refresh_token"`
}

type ToodledoConfigDatabase struct {
	DataFile string `mapstructure:"data_file"`
}

type ToodledoConfig struct {
	Auth     ToodledoConfigAuth     `mapstructure:"auth"`
	Database ToodledoConfigDatabase `mapstructure:"database"`
}

func NewToodledoConfig(configs Configs) ToodledoConfig {
	// TODO dirty, but it works
	return *configs.Get()
}
