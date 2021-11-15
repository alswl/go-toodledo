package itinjector

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/app"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dao"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/google/wire"
)

var IntegrationTestSet = wire.NewSet(
	common.NewConfigsForTesting,
	common.NewToodledoConfig,
	dao.NewBoltDB,
	client.NewToodledoCli,
	client.NewAuthFromConfigs,
	client.NewOAuth2ConfigFromConfigs,
	services.NewAccountService,
	services.NewTaskService,
	services.NewFolderService,
	services.NewFolderCachedService,
	app.NewToodledoCliApp,
)
