package app

import (
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
)

type ToodledoCliApp struct {
	// TODO service here
	Auth      runtime.ClientAuthInfoWriter
	TaskSvc   services.TaskService
	FolderSvc services.FolderService
}

func NewToodledoCliApp(auth runtime.ClientAuthInfoWriter,
	taskSvc services.TaskService,
	folderSvc services.FolderService,
) *ToodledoCliApp {
	return &ToodledoCliApp{Auth: auth, TaskSvc: taskSvc, FolderSvc: folderSvc}
}
