package fetchers

import (
	"context"
	"fmt"
	"time"

	"github.com/alswl/go-toodledo/pkg/common"

	"github.com/sirupsen/logrus"
)

const defaultFetchQueueSize = 1

// DaemonFetcher is a interface for one fetcher
// it runs in background and fetch data from remote.
type DaemonFetcher interface {
	Start(context.Context)
	// Stop the fetcher
	// TODO using ch to stop
	Stop()
	// Notify is used to notify fetcher to fetch data now.
	// isHardRefresh is used to indicate whether to clean cache.
	// return a chan to notify whether fetch success.
	Notify(isHardRefresh bool) (chan bool, error)
	Fetch(isHardRefresh bool) error
	// UIRefresh is used to notify ui app to refresh
	UIRefresh() chan bool
}

// FetchFn fetch data.
type FetchFn func(sd common.StatusDescriber, isHardRefresh bool) error

type FetchPromise struct {
	Done          chan bool
	IsHardRefresh bool
}

type intervalDaemonFetcher struct {
	ticker   *time.Ticker
	stop     chan struct{}
	fetchNow chan *FetchPromise
	// fetchForceNow chan bool
	uiRefresh chan bool
	// refreshed chan bool

	log             logrus.FieldLogger
	fn              FetchFn
	statusDescriber common.StatusDescriber
}

func NewSimpleFetcher(
	log logrus.FieldLogger,
	interval time.Duration,
	fn FetchFn,
	statusDescriber common.StatusDescriber,
) DaemonFetcher {
	var ticker *time.Ticker
	if interval == 0 {
		// do not run ticker
		ticker = time.NewTicker(1 * time.Second)
		ticker.Stop()
	} else {
		ticker = time.NewTicker(interval)
	}
	return &intervalDaemonFetcher{
		ticker:          ticker,
		stop:            make(chan struct{}),
		log:             log,
		fn:              fn,
		fetchNow:        make(chan *FetchPromise, defaultFetchQueueSize),
		statusDescriber: statusDescriber,
		uiRefresh:       make(chan bool),
		//refreshed:       make(chan bool, defaultRefreshQueueSize),
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
		case promise := <-s.fetchNow:
			s.log.WithField("ifHardRefresh", promise.IsHardRefresh).Info("fetcher now")
			err := s.Fetch(promise.IsHardRefresh)
			if err != nil {
				s.log.Errorf("fetcher fetch error: %v", err)
				promise.Done <- false
				// s.refreshed <- false
			}
			go func() {
				// TODO maybe leaks
				promise.Done <- true
			}()
			// s.refreshed <- true
			s.log.WithField("isHardRefresh", promise.IsHardRefresh).Info("fetcher now done")
		case <-s.ticker.C:
			s.log.Info("fetcher tick")
			err := s.Fetch(false)
			if err != nil {
				s.log.Errorf("fetcher fetch error: %v", err)
			}
			s.log.Info("fetcher tick done")
			s.UIRefresh() <- false
		}
	}
}

func (s *intervalDaemonFetcher) Start(ctx context.Context) {
	s.log.Info("fetcher started")
	go s.run(ctx)
}

// Fetch is used to fetch data from remote
// it was synchronized.
func (s *intervalDaemonFetcher) Fetch(hardRefresh bool) error {
	return s.fn(s.statusDescriber, hardRefresh)
}

func (s *intervalDaemonFetcher) Stop() {
	s.ticker.Stop()
	close(s.stop)
	s.log.Info("fetcher stopped")
}

// Notify is used to notify fetcher to fetch data now.
func (s *intervalDaemonFetcher) Notify(isHardRefresh bool) (chan bool, error) {
	s.log.WithField("isHardRefresh", isHardRefresh).Info("Notify")
	promise := &FetchPromise{
		Done:          make(chan bool),
		IsHardRefresh: isHardRefresh,
	}

	select {
	case s.fetchNow <- promise:
		return promise.Done, nil
	default:
		s.log.Info("fetcher is busy")
		return nil, fmt.Errorf("fetcher is busy")
	}
	// TODO using specific channel for fetch
	// return s.refreshed, nil
}

func (s *intervalDaemonFetcher) UIRefresh() chan bool {
	return s.uiRefresh
}
