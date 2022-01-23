package syncer

import (
	"context"
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
}

func NewToodledoSyncer(folderSvc services.FolderCachedService, accountSvc services.AccountService) ToodledoSyncer {
	ts := toodledoSyncer{
		log:        logrus.New(),
		folderSvc:  folderSvc,
		accountSvc: accountSvc,
	}
	syncer := NewSimpleSyncer(1*time.Minute, ts.sync)
	ts.Syncer = syncer
	return &ts
}

func (s *toodledoSyncer) SyncOnce() error {
	return s.sync()
}

func (s *toodledoSyncer) sync() error {
	me, err := s.accountSvc.Me()
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

	return nil
}
