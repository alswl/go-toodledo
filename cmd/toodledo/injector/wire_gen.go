// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
)

// Injectors from injector.go:

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	clientAuthInfoWriter, err := client.ProvideSimpleAuth()
	if err != nil {
		return nil, err
	}
	return clientAuthInfoWriter, nil
}

func InitTaskService() (services.TaskService, error) {
	clientAuthInfoWriter, err := client.ProvideSimpleAuth()
	if err != nil {
		return nil, err
	}
	taskService := services.ProvideTaskService(clientAuthInfoWriter)
	return taskService, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	clientAuthInfoWriter, err := client.ProvideSimpleAuth()
	if err != nil {
		return nil, err
	}
	taskService := services.ProvideTaskService(clientAuthInfoWriter)
	toodledoCliApp := app.NewToodledoCliApp(clientAuthInfoWriter, taskService)
	return toodledoCliApp, nil
}
