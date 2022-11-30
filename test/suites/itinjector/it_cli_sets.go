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

var CLISet = wire.NewSet(
	common.NewCliConfigFromViper,
	common.NewConfigCliConfig,
	logging.ProvideLogger,
	models.NewDefaultToodledoConfigDatabase,

	dal.ProvideBackend,
	client.NewToodledo,
	//client.NewAuthFromViper,
	client.NewAuthFromConfig,
	client.NewOAuth2ConfigFromViper,

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
	services.NewTaskRichPersistenceService,

	// wire not support generic now
	//informers.ProvideTaskInformer,

	app.NewToodledoCLIApp,
)
