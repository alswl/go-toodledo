package common

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewConfigFromCliConfig(cfg ToodledoCliConfig) (ToodledoConfig, error) {
	return cfg.Auth, nil
}

// NewCliConfigFromViper build Configs from viper.
func NewCliConfigFromViper() (ToodledoCliConfig, error) {
	var conf = NewToodledoCliConfig()
	err := viper.Unmarshal(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func NewCliConfigForTesting() (ToodledoCliConfig, error) {
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	viper.AddConfigPath(path.Join(home, ".config", "toodledo"))
	viper.SetConfigName("toodledo-test")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return ToodledoCliConfig{}, fmt.Errorf("failed to read config: %w", err)
	}
	return NewCliConfigFromViper()
}

func NewCliConfigMockForTesting() (ToodledoCliConfig, error) {
	var conf = ToodledoCliConfig{
		Auth: ToodledoConfig{
			UserID:       "test-user-id",
			ClientID:     "test-client-id",
			ClientSecret: "test-client-secret",
			AccessToken:  "test-access-token",
			ExpiredAt:    "2099-11-26T01:27:20+08:00",
			RefreshToken: "test-refresh-token",
		},
		Database: ToodledoConfigDatabase{
			DataFile: "tmp.db",
			Buckets:  nil,
		},
		Environment:    map[string]*ToodledoConfigEnvironment{},
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
