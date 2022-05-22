package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/common"
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
	remoteSvc  *taskService
	accountSvc AccountService

	syncLock sync.Mutex
	db       dal.Backend
}

func NewTaskCachedService(remoteSvc *taskService, accountSvc AccountService, db dal.Backend) TaskCachedService {
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

func (s *taskCachedService) ListWithChanged(lastEditTime *int32, start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	return s.remoteSvc.ListWithChanged(lastEditTime, start, limit)
}

func (s *taskCachedService) ListDeleted(lastEditTime *int32) ([]*models.TaskDeleted, error) {
	return s.remoteSvc.ListDeleted(lastEditTime)
}

func (s *taskCachedService) syncWithFn(fnEdited func() ([]*models.Task, error), fnDeleted func() ([]*models.TaskDeleted, error)) error {
	editedTasks, err := fnEdited()
	if err != nil {
		return err
	}
	s.syncLock.Lock()
	defer s.syncLock.Unlock()

	if err != nil {
		return err
	}
	for _, f := range editedTasks {
		bytes, _ := json.Marshal(f)
		_ = s.db.Put(TaskBucket, strconv.Itoa(int(f.ID)), bytes)
	}

	tds, _ := fnDeleted()
	for _, td := range tds {
		_ = s.db.Remove(TaskBucket, strconv.Itoa(int(td.ID)))
	}
	return nil
}

func (s *taskCachedService) Sync() error {
	return s.syncWithFn(s.listAllRemote, func() ([]*models.TaskDeleted, error) {
		return []*models.TaskDeleted{}, nil
	})
}

func (s *taskCachedService) PartialSync(lastEditTime *int32) error {
	return s.syncWithFn(
		func() ([]*models.Task, error) { return s.listChanged(lastEditTime) },
		func() ([]*models.TaskDeleted, error) { return s.ListDeleted(lastEditTime) },
	)
}

func (s *taskCachedService) FindById(id int64) (*models.Task, error) {
	all, _, err := s.ListAll()
	if err != nil {
		return nil, err
	}
	filterHeadOpt := funk.Head(funk.Filter(all, func(t *models.Task) bool {
		return t.ID == id
	}))
	if filterHeadOpt == nil {
		return nil, common.ErrNotFound
	}
	head := filterHeadOpt.(*models.Task)
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
		// TODO validate
		//ts = make([]*models.Task, 0)
	}
	return all, nil
}

func (s *taskCachedService) listChanged(lastEditTime *int32) ([]*models.Task, error) {
	var ts, all []*models.Task
	var err error
	var pagination *models.PaginatedInfo
	var start = int64(0)
	var limit = MaxNumPerRequest

	// TODO query from local data
	for {
		ts, pagination, err = s.remoteSvc.ListWithChanged(lastEditTime, start, limit)
		if err != nil {
			return nil, err
		}
		if len(ts) == 0 || pagination.Num == 0 {
			break
		}
		all = append(all, ts...)
		start = start + limit
		// validate
		//ts = make([]*models.Task, 0)
	}
	return all, nil
}

func (s *taskCachedService) ListAll() ([]*models.Task, int, error) {
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

func (s *taskCachedService) List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
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
	if query.FolderID != 0 {
		ts = funk.Filter(ts, func(t *models.Task) bool {
			return t.Folder == query.FolderID
		}).([]*models.Task)
	}
	if query.GoalID != 0 {
		ts = funk.Filter(ts, func(t *models.Task) bool {
			return t.Goal == query.GoalID
		}).([]*models.Task)
	}
	return ts, nil
}

func (s *taskCachedService) Create(title string) (*models.Task, error) {
	created, err := s.remoteSvc.Create(title)
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
