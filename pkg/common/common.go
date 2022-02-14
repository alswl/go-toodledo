package common

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/mitchellh/go-homedir"
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
	var conf models.ToodledoCliConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		return models.ToodledoCliConfig{}, err
	}
	conf.Database.Buckets = models.DefaultBuckets
	if conf.Database.DataFile == "" {
		dir, err := homedir.Dir()
		if err != nil {
			return models.ToodledoCliConfig{}, err
		}
		conf.Database.DataFile = path.Join(dir, ".toodledo", "data", "data.db")
	}
	return conf, nil
}

// NewCliConfigForTesting ...
func NewCliConfigForTesting() (models.ToodledoCliConfig, error) {
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	viper.AddConfigPath(path.Join(home, ".config"))
	viper.SetConfigName("toodledo-test")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
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
