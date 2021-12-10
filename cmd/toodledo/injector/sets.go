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
	common.NewCliConfigFromViper,
	common.NewConfigCliConfig,

	dao.NewBoltDB,
	client.NewToodledoCli,
	client.NewAuthFromViper,
	client.ProvideOAuth2ConfigFromViper,

	services.CurrentUser,
	services.NewAccountService,
	services.NewTaskService,
	services.NewFolderService,
	services.NewFolderCachedService,
	services.NewContextService,
	services.NewContextCachedService,
	services.NewGoalService,

	app.NewToodledoCliApp,
)
