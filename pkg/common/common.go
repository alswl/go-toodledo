package common

import (
	"github.com/spf13/viper"
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
	return &configs{}, nil
}

func (c *configs) Get() *ToodledoConfig {
	return c.conf
}
