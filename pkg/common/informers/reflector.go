package informers

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

const maxQueueSize = 100000

type T string
type U *int64

type Reflector interface {
	// ListNewer lists all the items in the queue and returns the lastSynced value
	// toodledo not support list and watch
	// only support before / after, it will retrieve changed items
	ListNewer() error
	Run(stop <-chan struct{})
	NotifyModified(interface{})

	HasSynced() bool
	LastSynced() U
	Chan() <-chan T
}

type reflector struct {
	log      logrus.FieldLogger
	duration time.Duration

	ticker *time.Ticker
	// TODO read it
	queue      chan T
	notify     chan interface{}
	lastSynced U
	cancelCtx  context.Context

	// TODO generic
	watcher func(ctx context.Context, lastSynced U) ([]T, U, error)
}

func NewReflector(
	log logrus.FieldLogger,
	duration time.Duration,
	watcher func(ctx context.Context, lastSynced U) ([]T, U, error),
) Reflector {
	return &reflector{
		log:        log,
		duration:   duration,
		ticker:     time.NewTicker(duration),
		queue:      make(chan T, maxQueueSize),
		notify:     make(chan interface{}),
		lastSynced: nil,
		watcher:    watcher,
		cancelCtx:  context.Background(),
	}
}

func (r *reflector) Run(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			r.log.Info("reflector stopped")
			r.cancelCtx.Done()
			return
		case <-r.ticker.C:
			r.log.Info("reflector ticked")
			err := r.ListNewer()
			if err != nil {
				r.log.WithError(err).Error("reflector list new")
			}
		case <-r.notify:
			r.log.Info("reflector notified")
			err := r.ListNewer()
			if err != nil {
				r.log.WithError(err).Error("reflector list new")
			}
		}
	}
}

func (r *reflector) ListNewer() error {
	list, lastSynced, err := r.watcher(r.cancelCtx, r.lastSynced)
	if err != nil {
		return err
	}
	for _, v := range list {
		r.queue <- v
	}
	r.lastSynced = lastSynced
	return nil
}

func (r *reflector) NotifyModified(i interface{}) {
	r.notify <- i
}

func (r *reflector) HasSynced() bool {
	return len(r.queue) == 0
}

func (r *reflector) LastSynced() U {
	return r.lastSynced
}

func (r *reflector) Chan() <-chan T {
	return r.queue
}
