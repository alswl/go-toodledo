package informers

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/syncer2"
	"github.com/thoas/go-funk"
	"sync"
)

var _ services.TaskService = (*TaskInformer)(nil)

var taskInformer *TaskInformer
var taskInformerOnce sync.Once

type TaskInformer struct {
	syncer   syncer2.Syncer2[models.Task]
	syncLock sync.Mutex
	db       dal.Backend
	queue    chan syncer2.SyncEvent
}

func NewTaskInformer(syncer syncer2.TaskSyncer, db dal.Backend) *TaskInformer {
	return &TaskInformer{
		syncer: syncer,
		db:     db,
		queue:  make(chan syncer2.SyncEvent, 100),
	}
}

func ProvideTaskInformer(db dal.Backend) *TaskInformer {
	taskInformerOnce.Do(func() {
		// FIXME syncer is nil
		taskInformer = NewTaskInformer(nil, db)
	})
	return taskInformer
}

var TaskBucket = "tasks2"

func (h *TaskInformer) ListAll() ([]*models.Task, int, error) {
	// XXX test
	all, err := h.db.List(TaskBucket)
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

func (h *TaskInformer) FindById(id int64) (*models.Task, error) {
	// XXX test
	all, _, err := h.ListAll()
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

func (h *TaskInformer) List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	// TODO test
	all, _, err := h.ListAll()
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

func (h *TaskInformer) Create(title string) (*models.Task, error) {
	// XXX test
	local := &models.Task{
		Title: title,
	}
	h.queue <- syncer2.SyncEvent{
		Type: syncer2.SyncItemTypeCreate,
		Item: &models.Task{
			Title: title,
		},
	}

	// TODO returned task is only local, id is 0
	return local, nil
}

func (h *TaskInformer) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
	// XXX test
	t := query.ToModel()
	h.queue <- syncer2.SyncEvent{
		Type: syncer2.SyncItemTypeCreate,
		Item: t,
	}
	// TODO returned task is only local, id is 0
	return t, nil
}

func (h *TaskInformer) Delete(id int64) error {
	// XXX test
	h.queue <- syncer2.SyncEvent{
		Type: syncer2.SyncItemTypeDelete,
		Item: &models.Task{
			ID: id,
		},
	}
	return nil
}

func (h *TaskInformer) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
	// XXX test
	for _, id := range ids {
		h.Delete(id)
	}
	return []int64{}, []*models.TaskDeleteItem{}, nil
}

func (h *TaskInformer) Edit(id int64, t *models.Task) (*models.Task, error) {
	// XXX test
	h.queue <- syncer2.SyncEvent{
		Type: syncer2.SyncItemTypeUpdate,
		Item: t,
	}
	return t, nil
}

func (h *TaskInformer) EditByQuery(query *queries.TaskEditQuery) (*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (h *TaskInformer) Complete(id int64) (*models.Task, error) {
	// XXX test
	t := &models.Task{
		ID:        id,
		Completed: 1,
	}
	h.queue <- syncer2.SyncEvent{
		Type: syncer2.SyncItemTypeUpdate,
		Item: t,
	}
	return t, nil
}

func (h *TaskInformer) UnComplete(id int64) (*models.Task, error) {
	// XXX test
	t := &models.Task{
		ID:        id,
		Completed: 0,
	}
	h.queue <- syncer2.SyncEvent{
		Type: syncer2.SyncItemTypeUpdate,
		Item: t,
	}
	return t, nil
}

func (h *TaskInformer) ListDeleted(lastEditTime *int32) ([]*models.TaskDeleted, error) {
	//TODO implement me
	panic("implement me")
}
