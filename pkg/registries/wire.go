package registries

import (
	"github.com/alswl/go-toodledo/pkg/service"
	"github.com/google/wire"
)

func InitializeFolderService() service.FolderService {
	wire.Build()
	// TODO using wire to inject
	return &service.FolderServiceImpl{}
}
