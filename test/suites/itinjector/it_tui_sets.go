package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client0"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var IntegrationTestTUISet = wire.NewSet(
	dal.ProvideBackend,
	client0.NewToodledoClient,
	// client.NewAuthFromConfig,
	client0.NewAuthFromConfig,
	client0.NewOAuth2ConfigFromConfigs,

	services.CurrentUser,
	services.NewAccountExtService,
	services.ProvideTaskLocalExtService,
	services.ProvideTaskLocalExtServiceIft,
	services.NewFolderCachedService,
	services.NewContextCachedService,
	services.NewGoalCachedService,
	services.NewTaskRichService,
	services.NewSettingService,

	// wire not support generic now
	// informers.ProvideTaskInformer,

	app.NewToodledoTUIApp,
)
