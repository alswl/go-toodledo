// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
)

// Injectors from itinjector.go:

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	clientAuthInfoWriter, err := client.ProvideSimpleAuth()
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

func InitTaskService() (services.TaskService, error) {
	toodledo := client.NewToodledoCli()
	clientAuthInfoWriter, err := client.ProvideSimpleAuth()
	if err != nil {
		return nil, err
	}
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	return taskService, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	clientAuthInfoWriter, err := client.ProvideSimpleAuth()
	if err != nil {
		return nil, err
	}
	toodledo := client.NewToodledoCli()
	taskService := services.NewTaskService(toodledo, clientAuthInfoWriter)
	toodledoCliApp := app.NewToodledoCliApp(clientAuthInfoWriter, taskService)
	return toodledoCliApp, nil
}
