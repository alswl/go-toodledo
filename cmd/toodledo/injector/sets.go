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

var SuperSet = wire.NewSet(
	common.NewCliConfigFromViper,
	common.NewConfigCliConfig,
	logging.ProvideLogger,

	dal.ProvideBackend,
	client.NewToodledo,
	//client.NewAuthFromViper,
	client.NewAuthFromConfig,
	client.NewOAuth2ConfigFromViper,

	services.CurrentUser,
	services.NewAccountService,
	services.NewTaskService0,
	services.NewTaskService,
	services.NewTaskLocalService,
	services.NewFolderService,
	services.NewFolderLocalService,
	services.NewContextService,
	services.NewContextLocalService,
	services.NewGoalService,
	services.NewGoalLocalService,
	services.NewSavedSearchService,
	services.NewTaskRichCachedService,

	// wire not support generic now
	//informers.ProvideTaskInformer,

	app.NewToodledoTUIApp,
)
