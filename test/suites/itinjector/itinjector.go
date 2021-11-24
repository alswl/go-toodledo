//go:build wireinject
// +build wireinject

package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
)

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	wire.Build(IntegrationTestSet)
	return &client.SimpleAuth{}, nil
}

func NewConfigs() (common.Configs, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func NewToodledoCli() (*client.Toodledo, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitFolderService() (services.FolderService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitFolderCachedService() (services.FolderCachedService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitTaskService() (services.TaskService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}
