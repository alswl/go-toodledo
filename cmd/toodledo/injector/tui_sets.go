package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var TUISet = wire.NewSet(
	common.NewCliConfigFromViper,
	common.NewConfigCliConfig,
	logging.ProvideLogger,
	common.NewToodledoConfigDatabaseFromToodledoCliConfig,

	dal.ProvideBackend,
	client.NewToodledo,
	// client.NewAuthFromViper,
	client.NewAuthFromConfig,
	client.NewOAuth2ConfigFromViper,

	services.CurrentUser,
	services.NewAccountService,
	services.NewTaskService,
	services.ProvideTaskLocalExtService,
	services.ProvideTaskLocalExtServiceIft,
	services.NewFolderService,
	services.ProvideFolderCachedService,
	services.NewContextService,
	services.ProvideContextCachedService,
	services.NewGoalService,
	services.NewGoalCachedService,
	services.NewSavedSearchService,
	services.NewTaskRichPersistenceService,

	// wire not support generic now
	// informers.ProvideTaskInformer,

	app.NewToodledoTUIApp,
)
