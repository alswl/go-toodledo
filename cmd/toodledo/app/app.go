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
	AccountSvc     services.AccountService
	TaskSvc        services.TaskService
	TaskExtSvc     services.TaskPersistenceExtService
	FolderSvc      services.FolderService
	FolderExtSvc   services.FolderPersistenceService
	ContextSvc     services.ContextService
	ContextExtSvc  services.ContextPersistenceService
	GoalSvc        services.GoalService
	GoalExtSvc     services.GoalPersistenceService
	SavedSearchSvc services.SavedSearchService

	TaskRichSvc services.TaskRichService
	fetcher     fetchers.DaemonFetcher
}

func NewToodledoTUIApp(
	accountSvc services.AccountService,
	taskSvc services.TaskService,
	taskExtSvc services.TaskPersistenceExtService,
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
		AccountSvc:     accountSvc,
		TaskSvc:        taskSvc,
		TaskExtSvc:     taskExtSvc,
		FolderSvc:      folderSvc,
		FolderExtSvc:   folderCachedSvc,
		ContextSvc:     contextSvc,
		ContextExtSvc:  contextCachedSvc,
		GoalSvc:        goalSvc,
		GoalExtSvc:     goalCachedSvc,
		SavedSearchSvc: savedSearchSvc,
		TaskRichSvc:    taskRichSvc,
	}
}
