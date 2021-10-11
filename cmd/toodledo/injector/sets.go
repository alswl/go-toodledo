package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dao"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	common.NewConfigsFromViper,
	common.NewToodledoConfig,
	dao.NewBoltDB,
	client.NewToodledoCli,
	client.ProvideSimpleAuth,
	client.ProvideOAuth2Config,
	services.NewTaskService,
	services.NewFolderService,
	app.NewToodledoCliApp,
)
