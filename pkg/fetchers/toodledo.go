package fetchers

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
)

type ToodledoFetchFunc struct {
	log logrus.FieldLogger

	folderSvc  services.FolderPersistenceService
	contextSvc services.ContextPersistenceService
	goalSvc    services.GoalPersistenceService
	taskSvc    services.TaskPersistenceExtService
	accountSvc services.AccountService
}

func NewToodledoFetchFunc(
	log logrus.FieldLogger,
	folderSvc services.FolderPersistenceService,
	contextSvc services.ContextPersistenceService,
	goalSvc services.GoalPersistenceService,
	taskSvc services.TaskPersistenceExtService,
	accountSvc services.AccountService,
) *ToodledoFetchFunc {
	return &ToodledoFetchFunc{
		log:        log,
		folderSvc:  folderSvc,
		contextSvc: contextSvc,
		goalSvc:    goalSvc,
		taskSvc:    taskSvc,
		accountSvc: accountSvc,
	}
}

func NewToodledoFetchFnPartial(
	log logrus.FieldLogger,
	folderSvc services.FolderPersistenceService,
	contextSvc services.ContextPersistenceService,
	goalSvc services.GoalPersistenceService,
	taskSvc services.TaskPersistenceExtService,
	accountSvc services.AccountService,
) FetchFn {
	return NewToodledoFetchFunc(log, folderSvc, contextSvc, goalSvc, taskSvc, accountSvc).Fetch
}

func (s *ToodledoFetchFunc) Fetch(statusDescriber StatusDescriber) error {
	statusDescriber.Syncing()

	me, err := s.accountSvc.Me()
	if err != nil {
		statusDescriber.Error(fmt.Errorf("auth failed"))
		return err
	}
	lastFetchInfo, err := s.accountSvc.GetLastFetchInfo()
	if err != nil {
		statusDescriber.Error(fmt.Errorf("sync failed"))
		return err
	}
	if err != nil {
		statusDescriber.Error(fmt.Errorf("get user status failed"))
		s.log.WithError(err).Error("get user status failed")
		return err
	}

	if lastFetchInfo == nil || me.LasteditFolder > lastFetchInfo.LasteditFolder {
		s.log.Info("Fetching folders")
		err = s.folderSvc.Sync()
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch folders failed"))
			s.log.WithError(err).Error("fetch folders")
		}
	}
	if lastFetchInfo == nil || me.LasteditContext > lastFetchInfo.LasteditContext {
		s.log.Info("Fetching contexts")
		err = s.contextSvc.Sync()
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch contexts failed"))
			s.log.WithError(err).Error("fetch contexts")
		}
	}
	if lastFetchInfo == nil || me.LasteditGoal > lastFetchInfo.LasteditGoal {
		s.log.Info("Fetching goals")
		err = s.goalSvc.Sync()
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch goals failed"))
			s.log.WithError(err).Error("fetch goals")
		}
	}
	if lastFetchInfo == nil || me.LasteditTask > lastFetchInfo.LasteditTask {
		s.log.Info("Fetching tasks")
		var lastEditTime *int32
		if lastFetchInfo != nil {
			lastEditTime = &lastFetchInfo.LasteditTask
		}
		err = s.taskSvc.PartialSync(lastEditTime)
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch tasks failed"))
			s.log.WithError(err).Error("fetch tasks")
		}
	}

	err = s.accountSvc.SetLastFetchInfo(me)
	if err != nil {
		statusDescriber.Error(fmt.Errorf("set last fetch info failed"))
		s.log.WithError(err).Error("set last fetch info")
	}
	statusDescriber.Success()
	return nil
}
