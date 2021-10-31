package common

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type Configs interface {
	// TODO using this instead of viper
	Get() *ToodledoConfig
}

type configs struct {
	conf *ToodledoConfig
}

func NewConfigsFromViper() (Configs, error) {
	var conf ToodledoConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return &configs{&conf}, nil
}

func NewConfigsForTesting() (Configs, error) {
	path := os.Getenv("TOODLEDO_CONFIG")
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf ToodledoConfig
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}
	return &configs{&conf}, nil
}

func (c *configs) Get() *ToodledoConfig {
	return c.conf
}
