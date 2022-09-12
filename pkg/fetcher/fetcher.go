package fetcher

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Fetcher interface {
	Start(context.Context)
	// Stop the fetcher
	// TODO using ch to stop
	Stop()
	fetch() error
}

type simpleFetcherWrapper struct {
	ticker   *time.Ticker
	stop     chan struct{}
	fetchNow chan bool

	log *logrus.Logger
	fn  func() error
}

func NewSimpleFetcher(interval time.Duration, fn func() error, logger *logrus.Logger) Fetcher {
	stop := make(chan struct{})
	return &simpleFetcherWrapper{ticker: time.NewTicker(interval), stop: stop, log: logger, fn: fn}
}

func (s *simpleFetcherWrapper) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			s.log.Info("fetcher stopped, ctx done")
			s.Stop()
			break
		case <-s.stop:
			s.log.Info("fetcher stopped, stop chan")
			break
		case <-s.ticker.C:
			s.log.Info("fetcher tick")
			err := s.fetch()
			if err != nil {
				s.log.Errorf("fetcher fetch error: %v", err)
			}
		}
	}
}

func (s *simpleFetcherWrapper) Start(ctx context.Context) {
	go s.run(ctx)
}

func (s *simpleFetcherWrapper) fetch() error {
	return s.fn()
}

func (s *simpleFetcherWrapper) Stop() {
	s.ticker.Stop()
	close(s.stop)
}
