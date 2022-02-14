package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/thoas/go-funk"
	"strconv"
	"sync"
)

type TaskCachedService interface {
	Cached
	TaskService

	ListAllByQuery(query *queries.TaskListQuery) ([]*models.Task, error)
}

type taskCachedService struct {
	remoteSvc  TaskService
	accountSvc AccountService

	syncLock sync.Mutex
	db       dal.Backend
}

func NewTaskCachedService(remoteSvc TaskService, accountSvc AccountService, db dal.Backend) TaskCachedService {
	return &taskCachedService{remoteSvc: remoteSvc, accountSvc: accountSvc, db: db}
}

var TaskBucket = "tasks"

var MaxNumPerRequest = int64(1000)

func (s *taskCachedService) LocalClear() error {
	err := s.db.Truncate(TaskBucket)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskCachedService) Sync() error {
	all, err := s.listAllRemote()
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

func (s *taskCachedService) FindById(id int64) (*models.Task, error) {
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

func (s *taskCachedService) listAllRemote() ([]*models.Task, error) {
	var ts, all []*models.Task
	var err error
	var pagination *models.PaginatedInfo
	var start = int64(0)
	var limit = MaxNumPerRequest

	// TODO query from local data
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

func (s *taskCachedService) ListAll() ([]*models.Task, error) {
	all, err := s.db.List(TaskBucket)
	if err != nil {
		return nil, err
	}
	var ts []*models.Task
	for _, v := range all {
		var t models.Task
		err = json.Unmarshal(v, &t)
		if err != nil {
			return nil, err
		}
		ts = append(ts, &t)
	}
	return ts, nil
}

func (s *taskCachedService) List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	// TODO test
	all, err := s.ListAll()
	if err != nil {
		return nil, nil, err
	}
	if start > int64(len(all)) {
		return nil, nil, nil
	}
	if start+limit > int64(len(all)) {
		limit = int64(len(all)) - start
	}
	return all[start : start+limit], &models.PaginatedInfo{
		Num:   start + limit,
		Total: int64(len(all)),
	}, nil
}

func (s *taskCachedService) ListAllByQuery(query *queries.TaskListQuery) ([]*models.Task, error) {
	all, err := s.db.List(TaskBucket)
	if err != nil {
		return nil, err
	}
	var ts []*models.Task
	for _, v := range all {
		var t models.Task
		err = json.Unmarshal(v, &t)
		if err != nil {
			return nil, err
		}
		ts = append(ts, &t)
	}
	if query.Priority != nil {
		ts = funk.Filter(ts, func(t *models.Task) bool {
			return tpriority.PriorityValue2Type(t.Priority) == *query.Priority
		}).([]*models.Task)
	}
	if query.Status != nil {
		ts = funk.Filter(ts, func(t *models.Task) bool {
			return tstatus.StatusValue2Type(t.Status) == *query.Status
		}).([]*models.Task)
	}
	if query.ContextID != 0 {
		ts = funk.Filter(ts, func(t *models.Task) bool {
			return t.Context == query.ContextID
		}).([]*models.Task)
	}
	return ts, nil
}

func (s *taskCachedService) Create(name string) (*models.Task, error) {
	created, err := s.remoteSvc.Create(name)
	if err != nil {
		return nil, err
	}
	// FIXME, using reconcile instead of cleanup
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *taskCachedService) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
	created, err := s.remoteSvc.CreateByQuery(query)
	if err != nil {
		return nil, err
	}
	// FIXME, using reconcile instead of cleanup
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *taskCachedService) Delete(id int64) error {
	err := s.remoteSvc.Delete(id)
	if err != nil {
		return err
	}
	// FIXME, using reconcile instead of cleanup
	err = s.LocalClear()
	if err != nil {
		return err
	}
	return nil
}

func (s *taskCachedService) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
	batch, items, err := s.remoteSvc.DeleteBatch(ids)
	if err != nil {
		return nil, nil, err
	}
	// FIXME, using reconcile instead of cleanup
	err = s.LocalClear()
	if err != nil {
		return nil, nil, err
	}
	return batch, items, nil
}

func (s *taskCachedService) Edit(id int64, t *models.Task) (*models.Task, error) {
	edited, err := s.remoteSvc.Edit(id, t)
	if err != nil {
		return nil, err
	}
	// FIXME, using reconcile instead of cleanup
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return edited, nil
}

func (s *taskCachedService) Complete(id int64) (*models.Task, error) {
	completed, err := s.remoteSvc.Complete(id)
	if err != nil {
		return nil, err
	}
	// FIXME, using reconcile instead of cleanup
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return completed, nil
}

func (s *taskCachedService) UnComplete(id int64) (*models.Task, error) {
	unCompleted, err := s.remoteSvc.UnComplete(id)
	if err != nil {
		return nil, err
	}
	// FIXME, using reconcile instead of cleanup
	err = s.LocalClear()
	if err != nil {
		return nil, err
	}
	return unCompleted, nil
}
