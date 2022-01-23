package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer"
	"github.com/google/wire"
)

var IntegrationTestSet = wire.NewSet(
	common.NewCliConfigForTesting,
	common.NewConfigCliConfig,

	dal.ProvideBackend,
	client.NewToodledo,
	//client.NewAuthFromConfig,
	client.NewAuthFromConfig,
	client.NewOAuth2ConfigFromConfigs,

	services.CurrentUser,
	services.NewAccountService,
	services.NewTaskService,
	services.NewTaskCachedService,
	services.NewFolderService,
	services.NewFolderCachedService,
	services.NewContextService,
	services.NewContextCachedService,
	services.NewGoalService,
	services.NewSavedSearchService,

	syncer.NewToodledoSyncer,

	app.NewToodledoCliApp,
)
