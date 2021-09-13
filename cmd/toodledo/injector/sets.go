package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	client.ProvideSimpleAuth,
	client.ProvideOAuth2Config,
	services.ProvideTaskService,
	app.NewToodledoCliApp,
	common.NewConfigsFromViper,
)
