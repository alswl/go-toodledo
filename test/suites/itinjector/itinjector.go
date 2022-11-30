//go:build wireinject
// +build wireinject

package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/google/wire"
)

func InitCLIBackend() (dal.Backend, error) {
	wire.Build(CLISet)
	return nil, nil
}

func InitCLIOption() (models.ToodledoCliConfig, error) {
	wire.Build(CLISet)
	return models.ToodledoCliConfig{}, nil
}

func InitCLIApp() (*app.ToodledoCLIApp, error) {
	wire.Build(CLISet)
	return nil, nil
}

func InitTUIApp() (*app.ToodledoTUIApp, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}
