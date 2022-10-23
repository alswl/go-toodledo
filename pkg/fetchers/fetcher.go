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
	Notify(isHardRefresh bool) error
	// UIRefresh is used to notify ui app to refresh
	UIRefresh() chan bool
}

// FetchFn fetch data
type FetchFn func(sd StatusDescriber, isHardRefresh bool) error

type intervalDaemonFetcher struct {
	ticker        *time.Ticker
	stop          chan struct{}
	fetchNow      chan bool
	fetchForceNow chan bool
	uiRefresh     chan bool

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
		uiRefresh:       make(chan bool),
	}
}

func (s *intervalDaemonFetcher) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			s.log.Info("fetcher stopped, ctx done")
			s.Stop()
			s.UIRefresh() <- true
			break
		case <-s.stop:
			s.log.Info("fetcher stopped, stop chan")
			s.UIRefresh() <- true
			break
		case isHardRefresh := <-s.fetchNow:
			s.log.WithField("ifHardRefresh", isHardRefresh).Info("fetcher now")
			err := s.fetch(isHardRefresh)
			if err != nil {
				s.log.Errorf("fetcher fetch error: %v", err)
			}
			s.UIRefresh() <- true
		case <-s.ticker.C:
			s.log.Info("fetcher tick")
			err := s.fetch(false)
			if err != nil {
				s.log.Errorf("fetcher fetch error: %v", err)
			}
			s.UIRefresh() <- true
		}
	}
}

func (s *intervalDaemonFetcher) Start(ctx context.Context) {
	s.log.Info("fetcher started")
	go s.run(ctx)
}

// fetch is used to fetch data from remote
// it was synchronized
func (s *intervalDaemonFetcher) fetch(hardRefresh bool) error {
	return s.fn(s.statusDescriber, hardRefresh)
}

func (s *intervalDaemonFetcher) Stop() {
	s.ticker.Stop()
	close(s.stop)
	s.log.Info("fetcher stopped")
}

// Notify is used to notify fetcher to fetch data now
func (s *intervalDaemonFetcher) Notify(isHardRefresh bool) error {
	s.log.Info("fetcher notify")
	go func() {
		s.fetchNow <- isHardRefresh
	}()
	return nil
}

func (s *intervalDaemonFetcher) UIRefresh() chan bool {
	return s.uiRefresh
}
