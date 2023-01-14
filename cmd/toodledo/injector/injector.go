//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/google/wire"
)

//func InitSimpleService() (*services.SimpleServices, error) {
//	wire.Build(CLISet, TUISet)
//	return nil, nil
//}

func InitCLIOption() (common.ToodledoCliConfig, error) {
	wire.Build(CommonSet)
	return common.ToodledoCliConfig{}, nil
}

func InitCLIApp() (*app.ToodledoCLIApp, error) {
	wire.Build(CLISet, CommonSet)
	return nil, nil
}

func InitTUIApp() (*app.ToodledoTUIApp, error) {
	wire.Build(TUISet, CommonSet)
	return nil, nil
}
