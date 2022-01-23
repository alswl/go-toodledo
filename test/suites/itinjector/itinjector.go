//go:build wireinject
// +build wireinject

package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
)

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	wire.Build(IntegrationTestSet)
	return &client.SimpleAuth{}, nil
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

func InitContextService() (services.ContextService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitContextCachedService() (services.ContextCachedService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitTaskService() (services.TaskService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitGoalsService() (services.GoalService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitAccountSvc() (services.AccountService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitSavedSearchService() (services.SavedSearchService, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitSyncer() (syncer.ToodledoSyncer, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	wire.Build(IntegrationTestSet)
	return nil, nil
}
