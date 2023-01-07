package app

import (
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/services"
)

type ToodledoCLIApp struct {
	// TODO move to services
	AccountSvc     services.AccountService
	TaskSvc        services.TaskService
	FolderSvc      services.FolderService
	ContextSvc     services.ContextService
	GoalSvc        services.GoalService
	SavedSearchSvc services.SavedSearchService

	TaskRichSvc services.TaskRichService

	Config common.ToodledoCliConfig
}

func NewToodledoCLIApp(
	accountSvc services.AccountService,
	taskSvc services.TaskService,
	folderSvc services.FolderService,
	contextSvc services.ContextService,
	goalSvc services.GoalService,
	savedSearchSvc services.SavedSearchService,
	taskRichSvc services.TaskRichService,
	config common.ToodledoCliConfig,
) *ToodledoCLIApp {
	return &ToodledoCLIApp{
		AccountSvc:     accountSvc,
		TaskSvc:        taskSvc,
		FolderSvc:      folderSvc,
		ContextSvc:     contextSvc,
		GoalSvc:        goalSvc,
		SavedSearchSvc: savedSearchSvc,
		TaskRichSvc:    taskRichSvc,
		Config:         config,
	}
}

type ToodledoTUIApp struct {
	// TODO move to services
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
	SettingSvc     services.SettingService

	TaskRichSvc services.TaskRichService
	// fetcher     fetchers.DaemonFetcher
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
	settingSvc services.SettingService,
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
		SettingSvc:     settingSvc,
	}
}
