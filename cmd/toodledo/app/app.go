package app

import (
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
)

type ToodledoCliApp struct {
	// TODO service here
	Auth       runtime.ClientAuthInfoWriter
	TaskSvc    services.TaskService
	FolderSvc  services.FolderCachedService
	ContextSvc services.ContextCachedService
	AccountSvc services.AccountService
	// TODO add Current account
}

func NewToodledoCliApp(auth runtime.ClientAuthInfoWriter, taskSvc services.TaskService, folderSvc services.FolderCachedService, contextSvc services.ContextCachedService, accountSvc services.AccountService) *ToodledoCliApp {
	return &ToodledoCliApp{Auth: auth, TaskSvc: taskSvc, FolderSvc: folderSvc, ContextSvc: contextSvc, AccountSvc: accountSvc}
}
