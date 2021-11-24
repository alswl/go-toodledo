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
	clientAuthInfoWriter, err := client.NewAuthFromViper()
	if err != nil {
		return nil, err
	}
	return clientAuthInfoWriter, nil
}

func NewConfigs() (common.Configs, error) {
	configs, err := common.NewConfigsFromViper()
	if err != nil {
		return nil, err
	}
	return configs, nil
}

func NewToodledoCli() (*client.Toodledo, error) {
	toodledo := client.NewToodledoCli()
	return toodledo, nil
}

func InitFolderService() (services.FolderService, error) {
	toodledo := client.NewToodledoCli()
	clientAuthInfoWriter, err := client.NewAuthFromViper()
	if err != nil {
		return nil, err
	}
	folderService := services.NewFolderService(toodledo, clientAuthInfoWriter)
	return folderService, nil
}

func InitFolderCachedService() (services.FolderCachedService, error) {
	toodledo := client.NewToodledoCli()
	clientAuthInfoWriter, err := client.NewAuthFromViper()
	if err != nil {
		return nil, err
	}
	folderService := services.NewFolderService(toodledo, clientAuthInfoWriter)
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter)
	configs, err := common.NewConfigsFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig := common.NewToodledoConfig(configs)
	backend, err := dao.NewBoltDB(toodledoConfig)
	if err != nil {
		return nil, err
	}
	folderCachedService := services.NewFolderCachedService(folderService, accountService, backend)
	return folderCachedService, nil
}

func InitTaskService() (services.TaskService, error) {
	toodledo := client.NewToodledoCli()
	clientAuthInfoWriter, err := client.NewAuthFromViper()
	if err != nil {
		return nil, err
	}
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	return taskService, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	clientAuthInfoWriter, err := client.NewAuthFromViper()
	if err != nil {
		return nil, err
	}
	toodledo := client.NewToodledoCli()
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	folderService := services.NewFolderService(toodledo, clientAuthInfoWriter)
	accountService := services.NewAccountService(toodledo, clientAuthInfoWriter)
	configs, err := common.NewConfigsFromViper()
	if err != nil {
		return nil, err
	}
	toodledoConfig := common.NewToodledoConfig(configs)
	backend, err := dao.NewBoltDB(toodledoConfig)
	if err != nil {
		return nil, err
	}
	folderCachedService := services.NewFolderCachedService(folderService, accountService, backend)
	toodledoCliApp := app.NewToodledoCliApp(clientAuthInfoWriter, taskService, folderCachedService, accountService)
	return toodledoCliApp, nil
}
