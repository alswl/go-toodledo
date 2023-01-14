package fetchers

import (
	"errors"
	"fmt"

	"github.com/alswl/go-toodledo/pkg/common"

	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
)

type ToodledoFetchFunc struct {
	log logrus.FieldLogger

	folderPstSvc  services.FolderPersistenceService
	contextPstSvc services.ContextPersistenceService
	goalPstSvc    services.GoalPersistenceService
	taskPstExtSvc services.TaskPersistenceExtService
	accountSvc    services.AccountExtService
}

func NewToodledoFetchFunc(
	log logrus.FieldLogger,
	folderPstSvc services.FolderPersistenceService,
	contextPstSvc services.ContextPersistenceService,
	goalPstSvc services.GoalPersistenceService,
	taskPstSvc services.TaskPersistenceExtService,
	accountSvc services.AccountExtService,
) *ToodledoFetchFunc {
	return &ToodledoFetchFunc{
		log:           log,
		folderPstSvc:  folderPstSvc,
		contextPstSvc: contextPstSvc,
		goalPstSvc:    goalPstSvc,
		taskPstExtSvc: taskPstSvc,
		accountSvc:    accountSvc,
	}
}

func NewToodledoFetchFnPartial(
	log logrus.FieldLogger,
	folderSvc services.FolderPersistenceService,
	contextSvc services.ContextPersistenceService,
	goalSvc services.GoalPersistenceService,
	taskSvc services.TaskPersistenceExtService,
	accountSvc services.AccountExtService,
) FetchFn {
	return NewToodledoFetchFunc(log, folderSvc, contextSvc, goalSvc, taskSvc, accountSvc).Fetch
}

func (s *ToodledoFetchFunc) Fetch(statusDescriber StatusDescriber, isHardRefresh bool) error {
	statusDescriber.Syncing()

	me, err := s.accountSvc.Me()
	if err != nil {
		statusDescriber.Error(fmt.Errorf("auth failed"))
		return err
	}
	lastFetchInfo, err := s.accountSvc.FindLastFetchInfo()
	if err != nil && !errors.Is(err, common.ErrNotFound) {
		statusDescriber.Error(fmt.Errorf("sync failed"))
		return err
	}

	if isHardRefresh || (lastFetchInfo == nil || me.LasteditFolder > lastFetchInfo.LasteditFolder) {
		s.log.Info("Fetching folders")
		err = s.folderPstSvc.Sync()
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch folders failed"))
			s.log.WithError(err).Error("fetch folders")
		}
	}
	if isHardRefresh || (lastFetchInfo == nil || me.LasteditContext > lastFetchInfo.LasteditContext) {
		s.log.Info("Fetching contexts")
		err = s.contextPstSvc.Sync()
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch contexts failed"))
			s.log.WithError(err).Error("fetch contexts")
		}
	}
	if isHardRefresh || (lastFetchInfo == nil || me.LasteditGoal > lastFetchInfo.LasteditGoal) {
		s.log.Info("Fetching goals")
		err = s.goalPstSvc.Sync()
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch goals failed"))
			s.log.WithError(err).Error("fetch goals")
		}
	}
	if isHardRefresh || (lastFetchInfo == nil || me.LasteditTask > lastFetchInfo.LasteditTask) {
		s.log.Info("Fetching tasks")
		var lastEditTime *int32
		if isHardRefresh {
			err = s.taskPstExtSvc.Clean()
			if err != nil {
				return err
			}
			err = s.taskPstExtSvc.Sync()
			if err != nil {
				return err
			}
		} else {
			if lastFetchInfo != nil {
				lastEditTime = &lastFetchInfo.LasteditTask
			}
			// TODO partial sync not works with nil(cache not clean, all is empty)
			err = s.taskPstExtSvc.PartialSync(lastEditTime)
		}
		if err != nil {
			statusDescriber.Error(fmt.Errorf("fetch tasks failed"))
			s.log.WithError(err).Error("fetch tasks")
		}
	}

	err = s.accountSvc.ModifyLastFetchInfo(me)
	if err != nil {
		statusDescriber.Error(fmt.Errorf("set last fetch info failed"))
		s.log.WithError(err).Error("set last fetch info")
	}
	statusDescriber.Success()
	return nil
}
