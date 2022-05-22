package syncer2

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

var _ Syncer2[models.Task] = (*syncer[models.Task])(nil)

type syncer[T models.Task] struct {
	ticker *time.Ticker
	stop   chan struct{}

	addFun    func()
	deleteFun func()
	updateFun func()
	queue     chan SyncEvent
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

func (s *syncer[T]) PostEvent(event SyncEvent) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) RegistryStatus(f func() (string, error)) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) add(event SyncEvent) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) delete(event SyncEvent) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) update(event SyncEvent) error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) syncFromRemote() error {
	//TODO implement me
	panic("implement me")
}

func (s *syncer[T]) syncToRemote(events []SyncEvent) error {
	//TODO implement me
	panic("implement me")
}

func NewSyncer() *syncer[models.Task] {
	s := &syncer[models.Task]{
		ticker: time.NewTicker(time.Minute * 5),
		queue:  make(chan SyncEvent, 100),
	}
	return s
}
