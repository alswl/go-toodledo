//go:build wireinject
// +build wireinject

package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
)

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	wire.Build(IntegrationTestSet)
	return &client.SimpleAuth{}, nil
}

func InitTaskService() (services.TaskService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}
