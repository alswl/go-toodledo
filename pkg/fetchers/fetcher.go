package fetchers

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

// DaemonFetcher is a interface for one fetcher
// it runs in background and fetch data from remote
type DaemonFetcher interface {
	Start(context.Context)
	// Stop the fetcher
	// TODO using ch to stop
	Stop()
	Notify() error
}

// FetchFn fetch data
type FetchFn func(sd StatusDescriber) error

type intervalDaemonFetcher struct {
	ticker   *time.Ticker
	stop     chan struct{}
	fetchNow chan bool

	log             logrus.FieldLogger
	fn              FetchFn
	statusDescriber StatusDescriber
}

func NewSimpleFetcher(log logrus.FieldLogger, interval time.Duration, fn FetchFn, statusDescriber StatusDescriber) DaemonFetcher {
	return &intervalDaemonFetcher{
		ticker:          time.NewTicker(interval),
		stop:            make(chan struct{}),
		log:             log,
		fn:              fn,
		fetchNow:        make(chan bool),
		statusDescriber: statusDescriber,
	}
}

func (s *intervalDaemonFetcher) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			s.log.Info("fetcher stopped, ctx done")
			s.Stop()
			break
		case <-s.stop:
			s.log.Info("fetcher stopped, stop chan")
			break
		case <-s.fetchNow:
			s.log.Info("fetcher now")
			err := s.fetch()
			if err != nil {
				s.log.Errorf("fetcher fetch error: %v", err)
			}
		case <-s.ticker.C:
			s.log.Info("fetcher tick")
			err := s.fetch()
			if err != nil {
				s.log.Errorf("fetcher fetch error: %v", err)
			}
		}
	}
}

func (s *intervalDaemonFetcher) Start(ctx context.Context) {
	s.log.Info("fetcher started")
	go s.run(ctx)
}

// fetch is used to fetch data from remote
// it was synchronized
func (s *intervalDaemonFetcher) fetch() error {
	return s.fn(s.statusDescriber)
}

func (s *intervalDaemonFetcher) Stop() {
	s.ticker.Stop()
	close(s.stop)
	s.log.Info("fetcher stopped")
}

// Notify is used to notify fetcher to fetch data now
func (s *intervalDaemonFetcher) Notify() error {
	s.log.Info("fetcher notify")
	go func() {
		s.fetchNow <- true
	}()
	return nil
}
