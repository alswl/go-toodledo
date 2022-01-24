package syncer

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
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
	log          *logrus.Logger
	lastSyncInfo *models.Account

	folderSvc  services.FolderCachedService
	accountSvc services.AccountService
	taskSvc    *services.TaskCachedService
	backend    dal.Backend
}

func NewToodledoSyncer(folderSvc services.FolderCachedService, accountSvc services.AccountService,
	taskSvc *services.TaskCachedService,
	backend dal.Backend) (ToodledoSyncer, error) {
	me, isCached, err := accountSvc.CachedMe()
	if err != nil {
		return nil, err
	}
	ts := toodledoSyncer{
		log:        logrus.New(),
		folderSvc:  folderSvc,
		accountSvc: accountSvc,
		taskSvc:    taskSvc,
	}
	// if it's found, using it from db
	// TODO better struct
	if isCached {
		ts.lastSyncInfo = me
	}
	syncer := NewSimpleSyncer(1*time.Minute, ts.sync)
	ts.Syncer = syncer
	return &ts, nil
}

func (s *toodledoSyncer) SyncOnce() error {
	return s.sync()
}

func (s *toodledoSyncer) sync() error {
	me, _, err := s.accountSvc.CachedMe()
	if err != nil {
		s.log.WithError(err).Error("Failed to get me in sync")
		return err
	}

	if s.lastSyncInfo == nil || me.LasteditFolder > s.lastSyncInfo.LasteditFolder {
		s.log.Info("Syncing folders")
		err = s.folderSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("Failed to sync folders")
		}
	}

	if s.lastSyncInfo == nil || me.LasteditTask > s.lastSyncInfo.LasteditTask {
		s.log.Info("Syncing tasks")
		err = s.taskSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("Failed to sync tasks")
		}
	}
	s.lastSyncInfo = me

	return nil
}
