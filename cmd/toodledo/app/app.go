package app

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer"
)

type ToodledoCliApp struct {
	CurrentUser *models.Account

	AccountSvc       services.AccountService
	TaskSvc          services.TaskService
	TaskCachedSvc    services.TaskCachedService
	FolderSvc        services.FolderService
	FolderCachedSvc  services.FolderCachedService
	ContextSvc       services.ContextService
	ContextCachedSvc services.ContextCachedService
	GoalSvc          services.GoalService
	GoalCachedSvc    services.GoalCachedService
	SavedSearchSvc   services.SavedSearchService

	TaskRichSvc services.TaskRichService
	Syncer      syncer.ToodledoFetcher
}

func NewToodledoCliApp(currentUser *models.Account, accountSvc services.AccountService, taskSvc services.TaskService, taskCachedSvc services.TaskCachedService, folderSvc services.FolderService, folderCachedSvc services.FolderCachedService, contextSvc services.ContextService, contextCachedSvc services.ContextCachedService, goalSvc services.GoalService, goalCachedSvc services.GoalCachedService, savedSearchSvc services.SavedSearchService, taskRichSvc services.TaskRichService, syncer syncer.ToodledoFetcher) *ToodledoCliApp {
	return &ToodledoCliApp{CurrentUser: currentUser, AccountSvc: accountSvc, TaskSvc: taskSvc, TaskCachedSvc: taskCachedSvc, FolderSvc: folderSvc, FolderCachedSvc: folderCachedSvc, ContextSvc: contextSvc, ContextCachedSvc: contextCachedSvc, GoalSvc: goalSvc, GoalCachedSvc: goalCachedSvc, SavedSearchSvc: savedSearchSvc, TaskRichSvc: taskRichSvc, Syncer: syncer}
}
