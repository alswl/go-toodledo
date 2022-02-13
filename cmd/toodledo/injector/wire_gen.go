// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer"
	"github.com/go-openapi/runtime"
)

// Injectors from injector.go:

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	return clientAuthInfoWriter, nil
}

func NewToodledoCli() (*client.Toodledo, error) {
	toodledo := client.NewToodledo()
	return toodledo, nil
}

func InitFolderService() (services.FolderService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	folderService := services.NewFolderService(toodledo, clientAuthInfoWriter)
	return folderService, nil
}

func InitFolderCachedService() (services.FolderCachedService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	folderService := services.NewFolderService(toodledo, clientAuthInfoWriter)
	backend, err := dal.ProvideBackend(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter, backend)
	folderCachedService := services.NewFolderCachedService(folderService, accountService, backend)
	return folderCachedService, nil
}

func InitContextService() (services.ContextService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	contextService := services.NewContextService(toodledo, clientAuthInfoWriter)
	return contextService, nil
}

func InitContextCachedService() (services.ContextCachedService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	contextService := services.NewContextService(toodledo, clientAuthInfoWriter)
	backend, err := dal.ProvideBackend(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter, backend)
	contextCachedService := services.NewContextCachedService(contextService, accountService, backend)
	return contextCachedService, nil
}

func InitTaskService() (services.TaskService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	return taskService, nil
}

func InitTaskCachedService() (*services.TaskCachedService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	backend, err := dal.ProvideBackend(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter, backend)
	taskCachedService := services.NewTaskCachedService(taskService, accountService, backend)
	return taskCachedService, nil
}

func InitGoalsService() (services.GoalService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	goalService := services.NewGoalService(toodledo, clientAuthInfoWriter)
	return goalService, nil
}

func InitAccountSvc() (services.AccountService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	backend, err := dal.ProvideBackend(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter, backend)
	return accountService, nil
}

func InitSavedSearchService() (services.SavedSearchService, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	savedSearchService := services.NewSavedSearchService(toodledo, clientAuthInfoWriter)
	return savedSearchService, nil
}

func InitSyncer() (syncer.ToodledoSyncer, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	folderService := services.NewFolderService(toodledo, clientAuthInfoWriter)
	backend, err := dal.ProvideBackend(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter, backend)
	folderCachedService := services.NewFolderCachedService(folderService, accountService, backend)
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	taskCachedService := services.NewTaskCachedService(taskService, accountService, backend)
	contextService := services.NewContextService(toodledo, clientAuthInfoWriter)
	contextCachedService := services.NewContextCachedService(contextService, accountService, backend)
	toodledoSyncer, err := syncer.NewToodledoSyncer(folderCachedService, accountService, taskCachedService, contextCachedService, backend)
	if err != nil {
		return nil, err
	}
	return toodledoSyncer, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	toodledo := client.NewToodledo()
	toodledoCliConfig, err := common.NewCliConfigFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig, err := common.NewConfigCliConfig(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := client.NewAuthFromConfig(toodledoConfig)
	if err != nil {
		return nil, err
	}
	backend, err := dal.ProvideBackend(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter, backend)
	account, err := services.CurrentUser(accountService)
	if err != nil {
		return nil, err
	}
	toodledoCliApp := app.NewToodledoCliApp(account)
	return toodledoCliApp, nil
}
