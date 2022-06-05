package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	common.NewCliConfigFromViper,
	common.NewConfigCliConfig,
	logging.ProvideLogger,
	logging.ProvideLoggerItf,

	dal.ProvideBackend,
	client.NewToodledo,
	//client.NewAuthFromViper,
	client.NewAuthFromConfig,
	client.NewOAuth2ConfigFromViper,

	services.CurrentUser,
	services.NewAccountService,
	services.NewTaskService0,
	services.NewTaskService,
	services.NewTaskCachedService,
	services.NewFolderService,
	services.NewFolderCachedService,
	services.NewContextService,
	services.NewContextCachedService,
	services.NewGoalService,
	services.NewGoalCachedService,
	services.NewSavedSearchService,

	// wire not support generic now
	//informers.ProvideTaskInformer,

	services.NewTaskRichService,

	syncer.NewToodledoSyncer,
	app.NewToodledoCliApp,
)
