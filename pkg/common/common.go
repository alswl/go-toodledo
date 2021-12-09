package common

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type Configs interface {
	// TODO using this instead of viper
	Get() *models.ToodledoCliConfig
}

// FIXME delete
type configs struct {
	conf *models.ToodledoCliConfig
}

// NewConfigsFromViper build Configs from viper
// XXX decoupling from viper
func NewConfigsFromViper() (Configs, error) {
	var conf models.ToodledoCliConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return &configs{&conf}, nil
}

func NewConfigsForTesting() (Configs, error) {
	path := os.Getenv("TOODLEDO_CONFIG")
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		path = home + "/" + ".toodledo-test.yaml"
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf models.ToodledoCliConfig
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}
	return &configs{&conf}, nil
}

func (c *configs) Get() *models.ToodledoCliConfig {
	return c.conf
}
