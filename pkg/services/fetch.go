package services

import (
	"errors"
	"fmt"

	"github.com/alswl/go-toodledo/pkg/fetchers"

	"github.com/alswl/go-toodledo/pkg/common"

	"github.com/sirupsen/logrus"
)

type ToodledoFetchService struct {
	log logrus.FieldLogger

	folderPstSvc  FolderPersistenceService
	contextPstSvc ContextPersistenceService
	goalPstSvc    GoalPersistenceService
	taskPstExtSvc TaskPersistenceExtService
	accountSvc    AccountExtService
}

func NewToodledoFetchService(
	log logrus.FieldLogger,
	folderPstSvc FolderPersistenceService,
	contextPstSvc ContextPersistenceService,
	goalPstSvc GoalPersistenceService,
	taskPstSvc TaskPersistenceExtService,
	accountSvc AccountExtService,
) *ToodledoFetchService {
	return &ToodledoFetchService{
		log:           log,
		folderPstSvc:  folderPstSvc,
		contextPstSvc: contextPstSvc,
		goalPstSvc:    goalPstSvc,
		taskPstExtSvc: taskPstSvc,
		accountSvc:    accountSvc,
	}
}

func NewToodledoFetchSvcsPartial(
	log logrus.FieldLogger,
	folderSvc FolderPersistenceService,
	contextSvc ContextPersistenceService,
	goalSvc GoalPersistenceService,
	taskSvc TaskPersistenceExtService,
	accountSvc AccountExtService,
) fetchers.FetchFn {
	return NewToodledoFetchService(log, folderSvc, contextSvc, goalSvc, taskSvc, accountSvc).Fetch
}

func (s *ToodledoFetchService) Fetch(statusDescriber common.StatusDescriber, isHardRefresh bool) error {
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
