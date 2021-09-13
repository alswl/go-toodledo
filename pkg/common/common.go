package common

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/spf13/viper"
)

type Configs interface {
	// TODO using this instead of viper
	Get() *models.ToodledoConfig
}

type configs struct {
	conf *models.ToodledoConfig
}

func NewConfigsFromViper() (Configs, error) {
	var conf models.ToodledoConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return &configs{&conf}, nil
}

func NewConfigsForTesting() (Configs, error) {
	return &configs{}, nil
}

func (c *configs) Get() *models.ToodledoConfig {
	return c.conf
}
