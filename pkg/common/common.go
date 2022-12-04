package common

import (
	"fmt"
	"os"
	"path"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewConfigCliConfig ...
func NewConfigCliConfig(cfg models.ToodledoCliConfig) (models.ToodledoConfig, error) {
	return cfg.Auth, nil
}

// NewCliConfigFromViper build Configs from viper.
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
	viper.SetConfigName("toodledo-test")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return models.ToodledoCliConfig{}, errors.Wrapf(err, "failed to read config")
	}
	return NewCliConfigFromViper()
}

// NewCliConfigMockForTesting ...
func NewCliConfigMockForTesting() (models.ToodledoCliConfig, error) {
	var conf = models.ToodledoCliConfig{
		Auth: models.ToodledoConfig{
			UserID:       "test-user-id",
			ClientID:     "test-client-id",
			ClientSecret: "test-client-secret",
			AccessToken:  "test-access-token",
			ExpiredAt:    "2099-11-26T01:27:20+08:00",
			RefreshToken: "test-refresh-token",
		},
		Database: models.ToodledoConfigDatabase{
			DataFile: "tmp.db",
			Buckets:  nil,
		},
		Environment:    map[string]*models.ToodledoConfigEnvironment{},
		DefaultContext: "default",
	}
	return conf, nil
}

// InitViper init viper with cfgFile or cfgDirInHome
// it must be called before viper.GetXXX, usually it was called in init().
func InitViper(cfgFile string, cfgDirInHome string, cfgName string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}

		// Search config in ~/.config/dir/conf.yaml
		viper.AddConfigPath(path.Join(home, cfgDirInHome))
		viper.SetConfigName("conf")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("config file", viper.ConfigFileUsed())
	}
}
