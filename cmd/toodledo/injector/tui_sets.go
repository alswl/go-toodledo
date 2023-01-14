package injector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var TUISet = wire.NewSet(
	services.CurrentUser,
	services.NewAccountExtService,

	services.ProvideTaskLocalExtService,
	services.ProvideTaskLocalExtServiceIft,
	services.ProvideFolderCachedService,
	services.ProvideContextCachedService,
	services.NewGoalCachedService,
	services.NewTaskRichPersistenceService,
	services.NewSettingService,

	// wire not support generic now
	// informers.ProvideTaskInformer,

	app.NewToodledoTUIApp,
)
