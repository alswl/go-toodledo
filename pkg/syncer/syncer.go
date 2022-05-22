package syncer

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Syncer interface {
	// TODO using ch to stop
	Start(context.Context)
	Stop()
	sync() error
}

type simpleSyncerWrapper struct {
	ticker  *time.Ticker
	stop    chan struct{}
	syncNow chan bool

	log *logrus.Logger
	fn  func() error
}

func NewSimpleSyncer(interval time.Duration, fn func() error, logger *logrus.Logger) Syncer {
	stop := make(chan struct{})
	return &simpleSyncerWrapper{ticker: time.NewTicker(interval), stop: stop, log: logger, fn: fn}
}

func (s *simpleSyncerWrapper) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			s.log.Info("syncer stopped, ctx done")
			s.Stop()
			break
		case <-s.stop:
			s.log.Info("syncer stopped, stop chan")
			break
		case <-s.ticker.C:
			s.log.Info("syncer tick")
			err := s.sync()
			if err != nil {
				s.log.Errorf("syncer sync error: %v", err)
			}
		}
	}
}

func (s *simpleSyncerWrapper) Start(ctx context.Context) {
	go s.run(ctx)
}

func (s *simpleSyncerWrapper) sync() error {
	return s.fn()
}

func (s *simpleSyncerWrapper) Stop() {
	s.ticker.Stop()
	close(s.stop)
}
