package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var IntegrationTestSet = wire.NewSet(
	auth.ProvideSimpleAuth,
	auth.ProvideOAuth2Config,
	services.ProvideTaskService,
	app.NewToodledoCliApp,
)
