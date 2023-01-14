package services

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/thoas/go-funk"
)

var instance TaskPersistenceExtService
var once sync.Once

const MaxNumPerRequest = int64(1000)

var TaskBucket = "tasks"

// var _ TaskPersistenceExtService = (*taskLocalExtService)(nil).
var _ TaskPersistenceExtService = (*taskLocalExtService)(nil)

var _ TaskExtendedService = (*taskLocalExtService)(nil)
var _ Synchronizable = (*taskLocalExtService)(nil)

type taskLocalExtService struct {
	taskSvc    TaskService
	accountSvc AccountExtService

	syncLock sync.Mutex
	db       dal.Backend
}

func NewTaskLocalExtService(
	taskSvc TaskService,
	accountSvc AccountExtService,
	db dal.Backend,
) TaskPersistenceExtService {
	return &taskLocalExtService{taskSvc: taskSvc, accountSvc: accountSvc, db: db}
}

func ProvideTaskLocalExtService(taskSvc TaskService, accountSvc AccountExtService,
	db dal.Backend) TaskPersistenceExtService {
	once.Do(func() {
		instance = NewTaskLocalExtService(taskSvc, accountSvc, db)
	})
	return instance
}

func ProvideTaskLocalExtServiceIft(
	taskSvc TaskService,
	accountSvc AccountExtService,
	db dal.Backend,
) TaskExtendedService {
	return ProvideTaskLocalExtService(taskSvc, accountSvc, db)
}

func (s *taskLocalExtService) Clean() error {
	err := s.db.Truncate(TaskBucket)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskLocalExtService) ListWithChanged(lastEditTime *int32, start, limit int64) (
	[]*models.Task, *models.PaginatedInfo, error) {
	return s.taskSvc.ListWithChanged(lastEditTime, start, limit)
}

func (s *taskLocalExtService) ListDeleted(lastEditTime *int32) ([]*models.TaskDeleted, error) {
	return s.taskSvc.ListDeleted(lastEditTime)
}

func (s *taskLocalExtService) syncWithFn(fnEdited func() ([]*models.Task, error), fnDeleted func() (
	[]*models.TaskDeleted, error)) error {
	editedTasks, err := fnEdited()
	if err != nil {
		return err
	}
	s.syncLock.Lock()
	defer s.syncLock.Unlock()

	for _, f := range editedTasks {
		// start := time.Now()
		bytes, _ := json.Marshal(f)
		_ = s.db.Put(TaskBucket, strconv.Itoa(int(f.ID)), bytes)
		// elapsed := time.Since(start)
		// logrus.WithField("elapsed", elapsed).WithField("title", f.Title).Info("syncWithFn")
	}

	tds, _ := fnDeleted()
	for _, td := range tds {
		_ = s.db.Remove(TaskBucket, strconv.Itoa(int(td.ID)))
	}
	return nil
}

func (s *taskLocalExtService) Sync() error {
	return s.syncWithFn(s.listAllRemote, func() ([]*models.TaskDeleted, error) {
		return []*models.TaskDeleted{}, nil
	})
}

func (s *taskLocalExtService) PartialSync(lastEditTime *int32) error {
	return s.syncWithFn(
		func() ([]*models.Task, error) { return s.listChanged(lastEditTime) },
		func() ([]*models.TaskDeleted, error) { return s.ListDeleted(lastEditTime) },
	)
}

func (s *taskLocalExtService) FindByID(id int64) (*models.Task, error) {
	bs, err := s.db.Get(TaskBucket, strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}
	var t models.Task
	err = json.Unmarshal(bs, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *taskLocalExtService) listAllRemote() ([]*models.Task, error) {
	var ts, all []*models.Task
	var err error
	var pagination *models.PaginatedInfo
	var start = int64(0)
	var limit = MaxNumPerRequest

	// TODO query from local data
	for {
		ts, pagination, err = s.taskSvc.List(start, limit)
		if err != nil {
			return nil, err
		}
		if len(ts) == 0 || pagination.Num == 0 {
			break
		}
		all = append(all, ts...)
		start += limit
		// TODO validate
		// ts = make([]*models.Task, 0)
	}
	return all, nil
}

func (s *taskLocalExtService) listChanged(lastEditTime *int32) ([]*models.Task, error) {
	var ts, all []*models.Task
	var err error
	var pagination *models.PaginatedInfo
	var start = int64(0)
	var limit = MaxNumPerRequest

	// TODO query from local data
	for {
		ts, pagination, err = s.taskSvc.ListWithChanged(lastEditTime, start, limit)
		if err != nil {
			return nil, err
		}
		if len(ts) == 0 || pagination.Num == 0 {
			break
		}
		all = append(all, ts...)
		start += limit
		// validate
		// ts = make([]*models.Task, 0)
	}
	return all, nil
}

// ListAll returns all tasks from cache, maybe cached missed
// FIXME avoid cache missed.
func (s *taskLocalExtService) ListAll() ([]*models.Task, int, error) {
	all, err := s.db.List(TaskBucket)
	if err != nil {
		return nil, 0, err
	}
	var ts []*models.Task
	for _, v := range all {
		var t models.Task
		err = json.Unmarshal(v, &t)
		if err != nil {
			return nil, 0, err
		}
		ts = append(ts, &t)
	}
	return ts, len(all), nil
}

func (s *taskLocalExtService) List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	// TODO test
	all, _, err := s.ListAll()
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

func (s *taskLocalExtService) ListAllByQuery(query *queries.TaskListQuery) ([]*models.Task, error) {
	// TODO refactor this
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
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return tpriority.Value2Type(t.Priority) == *query.Priority
		}).([]*models.Task)
	}
	if query.Status != nil {
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return tstatus.Value2Type(t.Status) == *query.Status
		}).([]*models.Task)
	}
	if query.ContextID == -1 {
		// None Context
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return t.Context == 0
		}).([]*models.Task)
	} else if query.ContextID != 0 {
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return t.Context == query.ContextID
		}).([]*models.Task)
	}
	if query.FolderID == -1 {
		// None Folder
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return t.Folder == 0
		}).([]*models.Task)
	} else if query.FolderID != 0 && query.FolderID != -1 {
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return t.Folder == query.FolderID
		}).([]*models.Task)
	}
	if query.GoalID == -1 {
		// None Goal
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return t.Goal == 0
		}).([]*models.Task)
	} else if query.GoalID != 0 {
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			return t.Goal == query.GoalID
		}).([]*models.Task)
	}
	if query.Incomplete != nil {
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			if *query.Incomplete {
				return t.Completed == 0
			}
			return t.Completed == 1
		}).([]*models.Task)
	} else {
		// nil Incomplete return incomplete + today complete
		ts, _ = funk.Filter(ts, func(t *models.Task) bool {
			if t.Completed == 0 {
				return true
			}
			now := time.Now()
			from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			to := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
			if t.Completed > from.Unix() && t.Completed < to.Unix() {
				return true
			}
			return false
		}).([]*models.Task)
	}

	return ts, nil
}

func (s *taskLocalExtService) Create(title string) (*models.Task, error) {
	return s.taskSvc.Create(title)
	// no cache clean, using fetcher to sync
}

func (s *taskLocalExtService) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
	return s.taskSvc.CreateByQuery(query)
	// no cache clean, using fetcher to sync
}

func (s *taskLocalExtService) Delete(id int64) error {
	return s.taskSvc.Delete(id)
	// no cache clean, using fetcher to sync
}

func (s *taskLocalExtService) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
	return s.taskSvc.DeleteBatch(ids)
	// no cache clean, using fetcher to sync
}

func (s *taskLocalExtService) Edit(id int64, t *models.TaskEdit) (*models.Task, error) {
	return s.taskSvc.Edit(id, t)
	// no cache clean, using fetcher to sync
}

func (s *taskLocalExtService) EditByQuery(_ *queries.TaskEditQuery) (*models.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (s *taskLocalExtService) Complete(id int64) (*models.Task, error) {
	return s.taskSvc.Complete(id)
	// no cache clean, using fetcher to partial sync
}

func (s *taskLocalExtService) UnComplete(id int64) (*models.Task, error) {
	return s.taskSvc.UnComplete(id)
	// no cache clean, using fetcher to partial sync
}

func (s *taskLocalExtService) Start(id int64) error {
	return s.taskSvc.Start(id)
}

func (s *taskLocalExtService) Stop(id int64) error {
	return s.taskSvc.Stop(id)
}
