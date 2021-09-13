//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
)

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	wire.Build(SuperSet)
	return &auth.SimpleAuth{}, nil
}

func InitTaskService() (services.TaskService, error) {
	wire.Build(SuperSet)
	return nil, nil
}

func InitApp() (*app.ToodledoCliApp, error) {
	wire.Build(SuperSet)
	return nil, nil
}
