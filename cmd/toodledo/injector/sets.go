package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	common.NewCliConfigFromViper,
	common.NewConfigCliConfig,

	dal.ProvideBackend,
	client.NewToodledo,
	//client.NewAuthFromViper,
	client.NewAuthFromConfig,
	client.ProvideOAuth2ConfigFromViper,

	services.CurrentUser,
	services.NewAccountService,
	services.NewTaskService,
	services.NewTaskCachedService,
	services.NewFolderService,
	services.NewFolderCachedService,
	services.NewContextService,
	services.NewContextCachedService,
	services.NewGoalService,
	services.NewGoalCachedService,
	services.NewSavedSearchService,

	services.NewTaskRichService,

	syncer.NewToodledoSyncer,

	app.NewToodledoCliApp,
)
