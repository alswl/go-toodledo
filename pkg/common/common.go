package common

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"
)

// NewConfigCliConfig ...
func NewConfigCliConfig(cfg models.ToodledoCliConfig) (models.ToodledoConfig, error) {
	return cfg.Auth, nil
}

// NewCliConfigFromViper build Configs from viper
func NewCliConfigFromViper() (models.ToodledoCliConfig, error) {
	var conf = models.NewToodledoCliConfig()
	err := viper.Unmarshal(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

// NewCliConfigForTesting ...
func NewCliConfigForTesting() (models.ToodledoCliConfig, error) {
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	viper.AddConfigPath(path.Join(home, ".config", "toodledo"))
	viper.SetConfigName(".toodledo-test")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return models.ToodledoCliConfig{}, errors.Wrapf(err, "failed to read config")
	}
	return NewCliConfigFromViper()
}

// NewCliConfigMockForTesting ...
func NewCliConfigMockForTesting() (models.ToodledoCliConfig, error) {
	var conf = models.ToodledoCliConfig{
		Auth:           models.ToodledoConfig{},
		Environment:    map[string]*models.ToodledoConfigEnvironment{},
		DefaultContext: "default",
		Database: models.ToodledoConfigDatabase{
			DataFile: "tmp.db",
			Buckets:  nil,
		},
	}
	return conf, nil
}
