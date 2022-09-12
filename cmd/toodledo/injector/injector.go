//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/fetcher"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
)

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	wire.Build(SuperSet)
	return &client.SimpleAuth{}, nil
}

func NewToodledoCli() (*client.Toodledo, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitFolderService() (services.FolderService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitFolderLocalService() (services.FolderLocalService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitContextService() (services.ContextService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitContextLocalService() (services.ContextLocalService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitTaskService() (services.TaskService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitTaskLocalService() (services.TaskLocalService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitGoalService() (services.GoalService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitGoalLocalService() (services.GoalLocalService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitAccountSvc() (services.AccountService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitSavedSearchService() (services.SavedSearchService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitTaskRichService() (services.TaskRichService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitTaskRichCachedService() (services.TaskRichCachedService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitSyncer() (fetcher.ToodledoFetcher, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitTUIApp() (*app.ToodledoTUIApp, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitCLIApp() (*app.ToodledoCLIApp, error) {
	wire.Build(CLISet)
	return nil, nil
}

//func InitTaskInformer() (*informers.TaskInformer, error) {
//	wire.Build(SuperSet)
//	return nil, nil
//}

func InitBackend() (dal.Backend, error) {
	wire.Build(SuperSet)
	return nil, nil
}
