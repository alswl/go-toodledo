package syncer

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"time"
)

type Progress interface {
	Start(int)
	Update(int)
	Finish()
}

type SyncItemType string
type SyncStatus string

var _ Syncer[models.Task] = (*syncer[models.Task])(nil)

type syncer[T models.Task] struct {
	ticker *time.Ticker
	stop   chan struct{}

	addFun    func()
	deleteFun func()
	updateFun func()
	queue     chan Event
}

func (s *syncer[T]) Status() string {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) SubscribeStatus(f func() (string, error)) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) AddFun(item T) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) DeleteFun(item T) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) UpdateFun(item T) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) Run(stopCh chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) PostEvent(event Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) RegistryStatus(f func() (string, error)) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) add(event Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) delete(event Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) update(event Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) syncFromRemote() error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) syncToRemote(events []Event) error {
	//TODO implement me
	panic("implement me")
}

func NewSyncer() *syncer[models.Task] {
	s := &syncer[models.Task]{
		ticker: time.NewTicker(time.Minute * 5),
		queue:  make(chan Event, 100),
	}
	return s
}
