package common

import "github.com/alswl/go-toodledo/pkg/models"

func NewToodledoConfig(configs Configs) models.ToodledoCliConfig {
	// TODO dirty, but it works
	cfg := *configs.Get()
	// TODO fix consts
	cfg.Database.Buckets = []string{
		"folders",
		"contexts",
	}
	return cfg
}
