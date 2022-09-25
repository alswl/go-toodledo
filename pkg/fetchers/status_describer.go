package fetchers

import (
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/sirupsen/logrus"
	"sync"
)

// StatusDescriber is a interface for describe the status of fetcher
// it gets progress from syncer, and invoke registered callback
type StatusDescriber interface {
	Progress() (int, int)
	SetProgress(current int, total int)

	OnSyncing(func() error)
	OnSuccess(func() error)
	OnError(func(err error) error)

	Error(err error)
	Success()
	Syncing()
}

var _ StatusDescriber = (*statusDescriber)(nil)

type statusDescriber struct {
	log     logrus.FieldLogger
	lock    sync.Mutex
	current int
	total   int
	err     error
	message string

	syncingFn func() error
	successFn func() error
	errorFn   func(err error) error
}

func NewStatusDescriber(onSyncing, onSuccess func() error, onError func(err error) error) StatusDescriber {
	return &statusDescriber{
		log:       logging.GetLogger("pkg.fetchers"),
		syncingFn: onSyncing,
		successFn: onSuccess,
		errorFn:   onError,
	}
}

func NewNoOpStatusDescriber() StatusDescriber {
	return &statusDescriber{
		log: logging.GetLogger("pkg.fetchers"),
		syncingFn: func() error {
			return nil
		},
		successFn: func() error {
			return nil
		},
		errorFn: func(err error) error { return err },
	}
}

func (s *statusDescriber) Success() {
	if s.successFn != nil {
		go func() {
			err := s.successFn()
			if err != nil {
				s.log.WithError(err).Error("successFn")
			}
		}()
	}
}

func (s *statusDescriber) Syncing() {
	if s.syncingFn != nil {
		go func() {
			err := s.syncingFn()
			if err != nil {
				s.log.WithError(err).Error("syncingFn")
			}
		}()
	}
}

func (s *statusDescriber) Error(err error) {
	if s.errorFn != nil {
		go func() {
			err := s.errorFn(err)
			if err != nil {
				s.log.WithError(err).Error("errorFn")
			}
		}()
	}
}

func (s *statusDescriber) OnSyncing(f func() error) {
	s.syncingFn = f
}

func (s *statusDescriber) OnSuccess(f func() error) {
	s.successFn = f
}

func (s *statusDescriber) OnError(f func(err error) error) {
	s.errorFn = f
}

func (s *statusDescriber) SetProgress(current int, total int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.current = current
	s.total = total
}

func (s *statusDescriber) Progress() (int, int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.current, s.total
}

type Progress func() int
