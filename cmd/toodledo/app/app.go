package app

import (
	"github.com/alswl/go-toodledo/pkg/fetchers"
	"github.com/alswl/go-toodledo/pkg/services"
)

type ToodledoCLIApp struct {
	AccountSvc     services.AccountService
	TaskSvc        services.TaskService
	FolderSvc      services.FolderService
	ContextSvc     services.ContextService
	GoalSvc        services.GoalService
	SavedSearchSvc services.SavedSearchService

	TaskRichSvc services.TaskRichService
}

func NewToodledoCLIApp(
	accountSvc services.AccountService,
	taskSvc services.TaskService,
	folderSvc services.FolderService,
	contextSvc services.ContextService,
	goalSvc services.GoalService,
	savedSearchSvc services.SavedSearchService,
	taskRichSvc services.TaskRichService,
) *ToodledoCLIApp {
	return &ToodledoCLIApp{
		AccountSvc:     accountSvc,
		TaskSvc:        taskSvc,
		FolderSvc:      folderSvc,
		ContextSvc:     contextSvc,
		GoalSvc:        goalSvc,
		SavedSearchSvc: savedSearchSvc,
		TaskRichSvc:    taskRichSvc,
	}
}

type ToodledoTUIApp struct {
	AccountSvc      services.AccountService
	TaskSvc         services.TaskService
	TaskLocalSvc    services.TaskPersistenceExtService
	FolderSvc       services.FolderService
	FolderLocalSvc  services.FolderPersistenceService
	ContextSvc      services.ContextService
	ContextLocalSvc services.ContextPersistenceService
	GoalSvc         services.GoalService
	GoalLocalSvc    services.GoalPersistenceService
	SavedSearchSvc  services.SavedSearchService

	TaskRichSvc services.TaskRichService
	fetcher     fetchers.DaemonFetcher
}

func NewToodledoTUIApp(
	accountSvc services.AccountService,
	taskSvc services.TaskService,
	taskCachedSvc services.TaskPersistenceExtService,
	folderSvc services.FolderService,
	folderCachedSvc services.FolderPersistenceService,
	contextSvc services.ContextService,
	contextCachedSvc services.ContextPersistenceService,
	goalSvc services.GoalService,
	goalCachedSvc services.GoalPersistenceService,
	savedSearchSvc services.SavedSearchService,
	taskRichSvc services.TaskRichService,
) *ToodledoTUIApp {
	return &ToodledoTUIApp{
		AccountSvc:      accountSvc,
		TaskSvc:         taskSvc,
		TaskLocalSvc:    taskCachedSvc,
		FolderSvc:       folderSvc,
		FolderLocalSvc:  folderCachedSvc,
		ContextSvc:      contextSvc,
		ContextLocalSvc: contextCachedSvc,
		GoalSvc:         goalSvc,
		GoalLocalSvc:    goalCachedSvc,
		SavedSearchSvc:  savedSearchSvc,
		TaskRichSvc:     taskRichSvc,
	}
}
