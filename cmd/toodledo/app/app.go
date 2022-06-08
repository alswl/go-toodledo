package app

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/services"
)

type ToodledoCliApp struct {
	CurrentUser *models.Account

	AccountSvc     services.AccountService
	TaskSvc        services.TaskService
	FolderSvc      services.FolderService
	ContextSvc     services.ContextService
	GoalSvc        services.GoalService
	SavedSearchSvc services.SavedSearchService
}

func NewToodledoCliApp(currentUser *models.Account, accountSvc services.AccountService, taskSvc services.TaskService, folderSvc services.FolderService, contextSvc services.ContextService, goalSvc services.GoalService, savedSearchSvc services.SavedSearchService) *ToodledoCliApp {
	return &ToodledoCliApp{CurrentUser: currentUser, AccountSvc: accountSvc, TaskSvc: taskSvc, FolderSvc: folderSvc, ContextSvc: contextSvc, GoalSvc: goalSvc, SavedSearchSvc: savedSearchSvc}
}
