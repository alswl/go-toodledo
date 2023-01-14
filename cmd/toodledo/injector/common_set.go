package injector

import (
	"github.com/alswl/go-toodledo/pkg/client0"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var CommonSet = wire.NewSet(
	common.NewCliConfigFromViper,
	common.NewConfigFromCliConfig,
	// TODO move to cli / tui
	client0.NewOAuth2ConfigFromViper,
	common.NewToodledoConfigDatabaseFromToodledoCliConfig,

	logging.ProvideLogger,
	dal.ProvideBackend,
	client0.NewToodledoClient,
	client0.NewAuthFromConfig,

	services.NewTaskService,
	services.NewFolderService,
	services.NewContextService,
	services.NewGoalService,
	services.NewSavedSearchService,
	services.NewAccountService,
)
