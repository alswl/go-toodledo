package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var CLISet = wire.NewSet(
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
