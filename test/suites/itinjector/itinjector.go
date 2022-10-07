//go:build wireinject
// +build wireinject

package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
)

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	wire.Build(IntegrationTestTUISet)
	return &client.SimpleAuth{}, nil
}

func NewToodledoCli() (*client.Toodledo, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitFolderService() (services.FolderService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitFolderLocalService() (services.FolderPersistenceService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitContextService() (services.ContextService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitContextLocalService() (services.ContextPersistenceService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitTaskService() (services.TaskService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitGoalService() (services.GoalService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitGoalLocalService() (services.GoalPersistenceService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitAccountSvc() (services.AccountService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitSavedSearchService() (services.SavedSearchService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitTaskRichService() (services.TaskRichService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitTaskLocalService() (services.TaskPersistenceExtService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitTaskExtendedService() (services.TaskPersistenceExtService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitTaskRichCachedService() (services.TaskRichService, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitTUIApp() (*app.ToodledoTUIApp, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}

func InitCLIApp() (*app.ToodledoCLIApp, error) {
	wire.Build(CLISet)
	return nil, nil
}

//func InitTaskInformer() (*informers.TaskInformer, error) {
//	wire.Build(IntegrationTestTUISet)
//	return nil, nil
//}

func InitBackend() (dal.Backend, error) {
	wire.Build(IntegrationTestTUISet)
	return nil, nil
}
