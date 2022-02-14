package syncer

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"time"
)

type ToodledoSyncer interface {
	Start(context.Context)
	Stop()
	sync() error
	SyncOnce() error
}

type toodledoSyncer struct {
	Syncer
	log *logrus.Logger

	folderSvc  services.FolderCachedService
	contextSvc services.ContextCachedService
	goalSvc    services.GoalCachedService
	taskSvc    services.TaskCachedService
	accountSvc services.AccountService
	backend    dal.Backend
}

func NewToodledoSyncer(
	folderSvc services.FolderCachedService,
	accountSvc services.AccountService,
	goalSvc services.GoalCachedService,
	taskSvc services.TaskCachedService,
	contextSvc services.ContextCachedService,
	backend dal.Backend) (ToodledoSyncer, error) {
	ts := toodledoSyncer{
		log:        logrus.New(),
		folderSvc:  folderSvc,
		contextSvc: contextSvc,
		goalSvc:    goalSvc,
		accountSvc: accountSvc,
		taskSvc:    taskSvc,
	}
	syncer := NewSimpleSyncer(1*time.Minute, ts.sync)
	ts.Syncer = syncer
	return &ts, nil
}

func (s *toodledoSyncer) SyncOnce() error {
	return s.sync()
}

func (s *toodledoSyncer) sync() error {
	me, err := s.accountSvc.Me()
	lastSyncInfo, err := s.accountSvc.GetLastSyncInfo()
	if err != nil {
		return err
	}
	if err != nil {
		s.log.WithError(err).Error("Failed to get me in sync")
		return err
	}

	if lastSyncInfo == nil || me.LasteditFolder > lastSyncInfo.LasteditFolder {
		s.log.Info("Syncing folders")
		err = s.folderSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("Failed to sync folders")
		}
	}
	if lastSyncInfo == nil || me.LasteditContext > lastSyncInfo.LasteditContext {
		s.log.Info("Syncing contexts")
		err = s.contextSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("Failed to sync contexts")
		}
	}
	if lastSyncInfo == nil || me.LasteditGoal > lastSyncInfo.LasteditGoal {
		s.log.Info("Syncing goals")
		err = s.goalSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("Failed to sync goals")
		}
	}
	if lastSyncInfo == nil || me.LasteditTask > lastSyncInfo.LasteditTask {
		s.log.Info("Syncing tasks")
		err = s.taskSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("Failed to sync tasks")
		}
	}

	err = s.accountSvc.SetLastSyncInfo(me)
	if err != nil {
		s.log.WithError(err).Error("Failed to set last sync info")
	}

	return nil
}
