package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client0"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var CLISet = wire.NewSet(
	dal.ProvideBackend,
	client0.NewToodledoClient,
	// client.NewAuthFromViper,
	client0.NewAuthFromConfig,
	client0.NewOAuth2ConfigFromViper,

	services.CurrentUser,
	services.NewAccountExtService,
	services.ProvideTaskLocalExtService,
	services.ProvideTaskLocalExtServiceIft,
	services.ProvideFolderCachedService,
	services.ProvideContextCachedService,
	services.NewGoalCachedService,
	services.NewTaskRichPersistenceService,

	// wire not support generic now
	// informers.ProvideTaskInformer,

	app.NewToodledoCLIApp,
)
