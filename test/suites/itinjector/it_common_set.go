package itinjector

import (
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var CommonSet = wire.NewSet(
	common.NewCliConfigMockForTesting,
	common.NewConfigFromCliConfig,

	logging.ProvideLogger,
	common.NewToodledoConfigDatabaseFromToodledoCliConfig,

	services.NewTaskService,
	services.NewFolderService,
	services.NewContextService,
	services.NewGoalService,
	services.NewSavedSearchService,
	services.NewAccountService,
)
