package services

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/alswl/go-toodledo/pkg/client0"
	"github.com/alswl/go-toodledo/pkg/client0/task"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	pointerutil "k8s.io/utils/pointer"
)

var DefaultFieldsInResponse = "folder,star,context,tag,goal,repeat,startdate,starttime,duedate,duetime,priority,length"

type TaskService interface {
	FindByID(id int64) (*models.Task, error)
	List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error)
	ListAll() ([]*models.Task, int, error)
	// TODO
	// ListAllAfter(after *time.Time) ([]*models.Task, error)
	// Create is simple create with only title, it was deprecated
	Create(title string) (*models.Task, error)
	CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error)
	Delete(id int64) error
	// DeleteBatch is batch delete tasks, return success ids, failed items and error
	DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error)
	Edit(id int64, t *models.TaskEdit) (*models.Task, error)
	EditByQuery(query *queries.TaskEditQuery) (*models.Task, error)
	Complete(id int64) (*models.Task, error)
	UnComplete(id int64) (*models.Task, error)
	ListDeleted(lastEditTime *int32) ([]*models.TaskDeleted, error)
	ListWithChanged(lastEditTime *int32, start, limit int64) ([]*models.Task, *models.PaginatedInfo, error)
	Start(id int64) error
	Stop(id int64) error
}

// TaskExtendedService is a service for tasks, it provided more query parameters.
type TaskExtendedService interface {
	TaskService

	// ListAllByQuery is list tasks by query, in-completed default, and with today complted tasks
	ListAllByQuery(query *queries.TaskListQuery) ([]*models.Task, error)
}

type TaskPersistenceExtService interface {
	TaskService
	TaskExtendedService
	Synchronizable
}

var _ TaskService = (*taskService)(nil)

// taskService is the implementation of TaskService by client.
type taskService struct {
	cli    *client0.Toodledo
	auth   runtime.ClientAuthInfoWriter
	logger logrus.FieldLogger
}

func NewTaskService(cli *client0.Toodledo, auth runtime.ClientAuthInfoWriter, logger logrus.FieldLogger) TaskService {
	return &taskService{cli: cli, auth: auth, logger: logger}
}

func (s *taskService) ListAll() ([]*models.Task, int, error) {
	logrus.Warnf("ListAll is not support, use List instead")
	return []*models.Task{}, 0, nil
}

func (s *taskService) FindByID(id int64) (*models.Task, error) {
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields())
	p.SetFields(&fields)
	p.SetID(&id)

	res, err := s.cli.Task.GetTasksGetPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	// TODO using multiple kind of payload item
	if len(res.Payload) == 1 {
		return nil, common.ErrNotFound
	}
	// _ := res.Payload[0].(models.PaginatedInfo)
	bytes, _ := json.Marshal(res.Payload[1])
	var t models.Task
	_ = json.Unmarshal(bytes, &t)
	return &t, nil
}

func (s *taskService) List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	return s.ListWithChanged(nil, start, limit)
}

func (s *taskService) ListWithChanged(lastEditTime *int32, start, limit int64) ([]*models.Task,
	*models.PaginatedInfo, error) {
	// TODO using TaskQuery,before, after,start,limit
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields())
	p.SetFields(&fields)
	comp := int64(-1)
	p.SetComp(&comp)
	num := limit
	p.SetNum(&num)
	startPtr := &start
	p.SetStart(startPtr)
	if lastEditTime != nil {
		p.After = utils.WrapPointerInt64((int64)(utils.UnwrapPointerInt32(lastEditTime)))
	}

	res, err := s.cli.Task.GetTasksGetPhp(p, s.auth)
	if err != nil {
		return nil, nil, err
	}
	var paging models.PaginatedInfo
	bytes, _ := json.Marshal(res.Payload[0])
	_ = json.Unmarshal(bytes, &paging)

	var tasks []*models.Task
	bytes, _ = json.Marshal(res.Payload[1:len(res.Payload)])
	_ = json.Unmarshal(bytes, &tasks)
	return tasks, &paging, nil
}

func (s *taskService) ListDeleted(lastEditTime *int32) ([]*models.TaskDeleted, error) {
	p := task.NewGetTasksDeletedPhpParams()
	if lastEditTime != nil {
		p.After = utils.WrapPointerInt64((int64)(utils.UnwrapPointerInt32(lastEditTime)))
	}

	res, err := s.cli.Task.GetTasksDeletedPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	var paging models.PaginatedInfo
	bytes, _ := json.Marshal(res.Payload[0])
	_ = json.Unmarshal(bytes, &paging)
	// fix official API missing total
	paging.Total = paging.Num

	var tds []*models.TaskDeleted
	bytes, _ = json.Marshal(res.Payload[1:len(res.Payload)])
	_ = json.Unmarshal(bytes, &tds)
	return tds, nil
}

func (s *taskService) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
	ts := []models.Task{*query.ToModel()}
	bytes, _ := json.Marshal(ts)
	bytesS := (string)(bytes)
	p := task.NewPostTasksAddPhpParams()
	p.Tasks = &bytesS
	p.Fields = &DefaultFieldsInResponse
	// https://api.toodledo.com/3/tasks/index.php#adding
	resp, err := s.cli.Task.PostTasksAddPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload[0], err
}

func (s *taskService) Create(title string) (*models.Task, error) {
	t := models.Task{
		Title: title,
	}
	bytes, _ := json.Marshal([]models.Task{t})
	bytesS := (string)(bytes)
	p := task.NewPostTasksAddPhpParams()
	p.Tasks = &bytesS
	p.Fields = pointerutil.String(enums.TaskFields2String(enums.GeneralTaskFields()))
	resp, err := s.cli.Task.PostTasksAddPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	if len(resp.Payload) == 0 {
		s.logger.WithField("title", title).Error("create task, return payload 0")
		return nil, fmt.Errorf("create task, return payload 0")
	}
	return resp.Payload[0], err
}

func (s *taskService) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
	p := task.NewPostTasksDeletePhpParams()
	idsString, _ := funk.Map(ids, func(x int64) string {
		return strconv.Itoa(int(x))
	}).([]string)
	bytes, _ := json.Marshal(idsString)
	bs := (string)(bytes)
	p.Tasks = &bs
	resp, err := s.cli.Task.PostTasksDeletePhp(p, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("delete batch request failed")
		return nil, nil, err
	}
	success, _ := funk.Filter(resp.Payload, func(x *models.TaskDeleteItem) bool {
		return x.Ref == ""
	}).([]*models.TaskDeleteItem)
	successIds, _ := funk.Map(success, func(x *models.TaskDeleteItem) int64 {
		return x.ID
	}).([]int64)

	failed, _ := funk.Filter(resp.Payload, func(x *models.TaskDeleteItem) bool {
		return x.Ref != ""
	}).([]*models.TaskDeleteItem)
	return successIds, failed, nil
}

func (s *taskService) Delete(id int64) error {
	t, err := s.FindByID(id)
	if err != nil {
		return err
	}
	success, failed, err := s.DeleteBatch([]int64{t.ID})
	if err != nil {
		return err
	}
	if len(failed) > 0 && len(success) == 1 {
		return fmt.Errorf("delete task %d", id)
	}
	return nil
}

func (s *taskService) Edit(id int64, t *models.TaskEdit) (*models.Task, error) {
	t.ID = id
	// FIXME @alswl
	// complete and timeon is not omitempty, so it will be initialized to 0
	// edit it will cause unexpected result
	bytes, _ := json.Marshal([]models.TaskEdit{*t})
	bytesS := (string)(bytes)
	p := task.NewPostTasksEditPhpParams()
	p.Tasks = &bytesS
	resp, err := s.cli.Task.PostTasksEditPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	if len(resp.Payload) == 0 {
		s.logger.WithField("id", id).WithField("resp", resp).Debug("edit task by query")
		return nil, fmt.Errorf("no response, edit task %d", id)
	}
	return resp.Payload[0], err
}

func (s *taskService) EditByQuery(query *queries.TaskEditQuery) (*models.Task, error) {
	ts := []models.Task{*query.ToModel()}
	bs, _ := json.Marshal(ts)
	bsStr := (string)(bs)
	p := task.NewPostTasksEditPhpParams()
	p.Tasks = &bsStr
	p.Fields = &DefaultFieldsInResponse
	// https://api.toodledo.com/3/tasks/index.php#editing
	resp, err := s.cli.Task.PostTasksEditPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload[0], err
}

func (s *taskService) Complete(id int64) (*models.Task, error) {
	_, err := s.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.Edit(id, &models.TaskEdit{
		Completed: utils.WrapPointerInt64(time.Now().Unix()),
		// auto reschedule by toodledo
		Reschedule: 1,
	})
}

func (s *taskService) UnComplete(id int64) (*models.Task, error) {
	_, err := s.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.Edit(id, &models.TaskEdit{
		Completed: utils.WrapPointerInt64(0),
	})
}

func (s *taskService) Start(id int64) error {
	t, err := s.FindByID(id)
	if err != nil {
		return err
	}
	if t.Timeron != 0 {
		return fmt.Errorf("task %d already started", id)
	}

	_, err = s.Edit(id, &models.TaskEdit{
		Timeron: utils.WrapPointerInt64(time.Now().Unix()),
	})
	return err
}

func (s *taskService) Stop(id int64) error {
	t, err := s.FindByID(id)
	if err != nil {
		return err
	}
	if t.Timeron == 0 {
		return fmt.Errorf("task %d not started", id)
	}

	_, err = s.Edit(id, &models.TaskEdit{
		Timer:   utils.WrapPointerInt64(t.Timer + time.Now().Unix() - t.Timeron),
		Timeron: utils.WrapPointerInt64(0),
	})
	return err
}

func rankTask(task *models.Task) int64 {
	var rank int64
	const scale1 = 10000
	const scale2 = 1000
	const scale3 = 100
	const scale100 = 100
	const to10 = 10.0
	const to11 = 11.0
	const maxTimestamp = 2147483647.0

	rank += int64((float64(task.Priority+1) / to10 * scale100) * scale1)
	statusValue := task.Status
	if statusValue == 0 {
		statusValue = 11
	}
	rank += int64(float64(statusValue)/to11*scale100) * scale2
	if task.Duedate != 0 {
		rank += int64(float64(maxTimestamp-task.Duedate) / maxTimestamp * scale100 * scale3)
	}
	return rank
}
func sortTasks(tasks []*models.Task) []*models.Task {
	// sort by rankTask
	sorted := make([]*models.Task, len(tasks))
	copy(sorted, tasks)
	sort.SliceStable(sorted, func(i, j int) bool {
		return rankTask(sorted[i]) > rankTask(sorted[j])
	})
	return sorted
}
