package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/thoas/go-funk"
	"strconv"
	"sync"
)

type TaskCachedService struct {
	remoteSvc  TaskService
	accountSvc AccountService

	syncLock *sync.Mutex
	db       dal.Backend
}

func NewTaskCachedService(remoteSvc TaskService, accountSvc AccountService, db dal.Backend) *TaskCachedService {
	return &TaskCachedService{remoteSvc: remoteSvc, accountSvc: accountSvc, db: db}
}

var TaskBucket = "tasks"

func (s *TaskCachedService) LocalClear() error {
	err := s.db.Truncate(TaskBucket)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskCachedService) Sync() error {
	// TODO using pagination
	all, err := s.ListAll()
	if err != nil {
		return err
	}
	s.syncLock.Lock()
	defer s.syncLock.Unlock()

	err = s.LocalClear()
	if err != nil {
		return err
	}
	for _, f := range all {
		bytes, _ := json.Marshal(f)
		s.db.Put(TaskBucket, strconv.Itoa(int(f.ID)), bytes)
	}
	return nil
}

func (s TaskCachedService) FindById(id int64) (*models.Task, error) {
	all, err := s.ListAll()
	if err != nil {
		return nil, err
	}
	filtered := funk.Filter(all, func(t *models.Task) bool {
		return t.ID == id
	}).([]*models.Task)
	head := funk.Head(filtered).(*models.Task)
	return head, nil
}

func (s TaskCachedService) ListAll() ([]*models.Task, error) {
	var ts, all []*models.Task
	var err error
	var pagination *models.PaginatedInfo
	var start = int64(0)
	var limit = int64(2)

	for {
		ts, pagination, err = s.remoteSvc.List(start, limit)
		if err != nil {
			return nil, err
		}
		if len(ts) == 0 || pagination.Num == 0 {
			break
		}
		all = append(all, ts...)
		start = start + limit
		ts = make([]*models.Task, 0)
	}
	return all, nil
}

func (s TaskCachedService) ListByQuery(query *queries.TaskSearchQuery) ([]*models.Task, *models.PaginatedInfo, error) {
	// FIXME api server did not support query, we must keep local storage.
	return []*models.Task{}, nil, nil
}

func (s TaskCachedService) Create(name string) (*models.Task, error) {
	created, err := s.remoteSvc.Create(name)
	if err != nil {
		return nil, err
	}
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s TaskCachedService) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
	created, err := s.remoteSvc.CreateByQuery(query)
	if err != nil {
		return nil, err
	}
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s TaskCachedService) Delete(id int64) error {
	err := s.remoteSvc.Delete(id)
	if err != nil {
		return err
	}
	err = s.LocalClear()
	if err != nil {
		return err
	}
	return nil
}

func (s TaskCachedService) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
	batch, items, err := s.remoteSvc.DeleteBatch(ids)
	if err != nil {
		return nil, nil, err
	}
	err = s.LocalClear()
	if err != nil {
		return nil, nil, err
	}
	return batch, items, nil
}

func (s TaskCachedService) Edit(id int64, t *models.Task) (*models.Task, error) {
	edited, err := s.remoteSvc.Edit(id, t)
	if err != nil {
		return nil, err
	}
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return edited, nil
}

func (s TaskCachedService) Complete(id int64) (*models.Task, error) {
	completed, err := s.remoteSvc.Complete(id)
	if err != nil {
		return nil, err
	}
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return completed, nil
}

func (s TaskCachedService) UnComplete(id int64) (*models.Task, error) {
	unCompleted, err := s.remoteSvc.UnComplete(id)
	if err != nil {
		return nil, err
	}
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return unCompleted, nil
}
