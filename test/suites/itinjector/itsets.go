package itinjector

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

var IntegrationTestSet = wire.NewSet(
	common.NewCliConfigForTesting,
	common.NewConfigCliConfig,
	logging.ProvideLogger,
	logging.ProvideLoggerItf,

	dal.ProvideBackend,
	client.NewToodledo,
	//client.NewAuthFromConfig,
	client.NewAuthFromConfig,
	client.NewOAuth2ConfigFromConfigs,

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
	services.NewTaskRichCachedService,

	// wire not support generic now
	//informers.ProvideTaskInformer,

	syncer.NewToodledoSyncer,
	app.NewToodledoCliApp,
)
