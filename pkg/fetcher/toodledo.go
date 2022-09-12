package fetcher

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"time"
)

type ToodledoFetcher interface {
	Start(context.Context)
	Stop()
	fetch() error
	FetchOnce() error
}

type toodledoFetcher struct {
	Fetcher
	log *logrus.Logger

	folderSvc  services.FolderLocalService
	contextSvc services.ContextLocalService
	goalSvc    services.GoalLocalService
	taskSvc    services.TaskLocalService
	accountSvc services.AccountService
	backend    dal.Backend
}

func NewToodledoFetcher(
	folderSvc services.FolderLocalService,
	accountSvc services.AccountService,
	goalSvc services.GoalLocalService,
	taskSvc services.TaskLocalService,
	contextSvc services.ContextLocalService,
	backend dal.Backend,
	logger *logrus.Logger) (ToodledoFetcher, error) {
	ts := toodledoFetcher{
		log:        logrus.New(),
		folderSvc:  folderSvc,
		contextSvc: contextSvc,
		goalSvc:    goalSvc,
		accountSvc: accountSvc,
		taskSvc:    taskSvc,
	}
	ts.Fetcher = NewSimpleFetcher(1*time.Minute, ts.fetch, logger)
	return &ts, nil
}

func (s *toodledoFetcher) FetchOnce() error {
	return s.fetch()
}

func (s *toodledoFetcher) fetch() error {
	me, _ := s.accountSvc.Me()
	lastFetchInfo, err := s.accountSvc.GetLastFetchInfo()
	if err != nil {
		return err
	}
	if err != nil {
		s.log.WithError(err).Error("get me in fetch")
		return err
	}

	if lastFetchInfo == nil || me.LasteditFolder > lastFetchInfo.LasteditFolder {
		s.log.Info("Fetching folders")
		err = s.folderSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("fetch folders")
		}
	}
	if lastFetchInfo == nil || me.LasteditContext > lastFetchInfo.LasteditContext {
		s.log.Info("Fetching contexts")
		err = s.contextSvc.Sync()
		if err != nil {
			s.log.WithError(err).Error("fetch contexts")
		}
	}
	if lastFetchInfo == nil || me.LasteditGoal > lastFetchInfo.LasteditGoal {
		s.log.Info("Fetching goals")
		err = s.goalSvc.Sync()
		if err != nil {
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
			s.log.WithError(err).Error("fetch tasks")
		}
	}

	err = s.accountSvc.SetLastFetchInfo(me)
	if err != nil {
		s.log.WithError(err).Error("set last fetch info")
	}

	return nil
}
