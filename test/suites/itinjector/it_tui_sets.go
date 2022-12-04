package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var IntegrationTestTUISet = wire.NewSet(
	common.NewCliConfigForTesting,
	common.NewConfigCliConfig,
	logging.ProvideLogger,
	models.NewToodledoConfigDatabaseFromToodledoCliConfig,

	dal.ProvideBackend,
	client.NewToodledo,
	// client.NewAuthFromConfig,
	client.NewAuthFromConfig,
	client.NewOAuth2ConfigFromConfigs,

	services.CurrentUser,
	services.NewAccountService,
	services.NewTaskService,
	services.ProvideTaskLocalExtService,
	services.ProvideTaskLocalExtServiceIft,
	services.NewFolderService,
	services.NewFolderCachedService,
	services.NewContextService,
	services.NewContextCachedService,
	services.NewGoalService,
	services.NewGoalCachedService,
	services.NewSavedSearchService,
	services.NewTaskRichService,

	// wire not support generic now
	// informers.ProvideTaskInformer,

	app.NewToodledoTUIApp,
)
