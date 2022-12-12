//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/google/wire"
)

func InitCLIBackend() (dal.Backend, error) {
	wire.Build(CLISet)
	return nil, nil
}

func InitCLIOption() (common.ToodledoCliConfig, error) {
	wire.Build(CLISet)
	return common.ToodledoCliConfig{}, nil
}

func InitCLIApp() (*app.ToodledoCLIApp, error) {
	wire.Build(CLISet)
	return nil, nil
}

func InitTUIApp() (*app.ToodledoTUIApp, error) {
	wire.Build(TUISet)
	return nil, nil
}
