// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dao"
	"github.com/alswl/go-toodledo/pkg/services"
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
	toodledo := client.NewToodledoCli()
	return toodledo, nil
}

func InitFolderService() (services.FolderService, error) {
	toodledo := client.NewToodledoCli()
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
	toodledo := client.NewToodledoCli()
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
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter)
	backend, err := dao.NewBoltDB(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	folderCachedService := services.NewFolderCachedService(folderService, accountService, backend)
	return folderCachedService, nil
}

func InitContextService() (services.ContextService, error) {
	toodledo := client.NewToodledoCli()
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
	toodledo := client.NewToodledoCli()
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
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter)
	backend, err := dao.NewBoltDB(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	contextCachedService := services.NewContextCachedService(contextService, accountService, backend)
	return contextCachedService, nil
}

func InitTaskService() (services.TaskService, error) {
	toodledo := client.NewToodledoCli()
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

func InitGoalsService() (services.GoalService, error) {
	toodledo := client.NewToodledoCli()
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

func InitApp() (*app.ToodledoCliApp, error) {
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
	toodledo := client.NewToodledoCli()
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	folderService := services.NewFolderService(toodledo, clientAuthInfoWriter)
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter)
	backend, err := dao.NewBoltDB(toodledoCliConfig)
	if err != nil {
		return nil, err
	}
	folderCachedService := services.NewFolderCachedService(folderService, accountService, backend)
	contextService := services.NewContextService(toodledo, clientAuthInfoWriter)
	contextCachedService := services.NewContextCachedService(contextService, accountService, backend)
	toodledoCliApp := app.NewToodledoCliApp(clientAuthInfoWriter, taskService, folderCachedService, contextCachedService, accountService)
	return toodledoCliApp, nil
}
