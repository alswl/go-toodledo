package services

import (
	"encoding/json"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/task"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	pointerutil "k8s.io/utils/pointer"
	"strconv"
	"time"
)

var DefaultFieldsInResponse = "folder,star,context,tag,goal,repeat,startdate,starttime,duedate,duetime,priority,length"

type TaskService interface {
	FindById(id int64) (*models.Task, error)
	List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error)
	ListAll() ([]*models.Task, int, error)
	// Create is simple create with only title, it was deprecated
	Create(title string) (*models.Task, error)
	CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error)
	Delete(id int64) error
	// DeleteBatch is batch delete tasks, return success ids, failed items and error
	DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error)
	Edit(id int64, t *models.Task) (*models.Task, error)
	EditByQuery(query *queries.TaskEditQuery) (*models.Task, error)
	Complete(id int64) (*models.Task, error)
	UnComplete(id int64) (*models.Task, error)
	ListDeleted(lastEditTime *int32) ([]*models.TaskDeleted, error)
}

type taskService struct {
	cli    *client.Toodledo
	auth   runtime.ClientAuthInfoWriter
	logger logrus.FieldLogger
}

func NewTaskService0(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter, logger logrus.FieldLogger) *taskService {
	return &taskService{cli: cli, auth: auth, logger: logger}
}

func NewTaskService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter, logger logrus.FieldLogger) TaskService {
	return NewTaskService0(cli, auth, logger)
}

func (s *taskService) ListAll() ([]*models.Task, int, error) {
	logrus.Warnf("ListAll is not support, use List instead")
	return []*models.Task{}, 0, nil
}

func (s *taskService) FindById(id int64) (*models.Task, error) {
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields)
	p.SetFields(&fields)
	p.SetID(&id)

	res, err := s.cli.Task.GetTasksGetPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	// TODO using multiple kind of payload item
	if err != nil || len(res.Payload) == 1 {
		return nil, common.ErrNotFound
	}
	//_ := res.Payload[0].(models.PaginatedInfo)
	bytes, _ := json.Marshal(res.Payload[1])
	var t models.Task
	_ = json.Unmarshal(bytes, &t)
	return &t, nil
}

func (s *taskService) List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	return s.ListWithChanged(nil, start, limit)
}

func (s *taskService) ListWithChanged(lastEditTime *int32, start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	// TODO using TaskQuery,before, after,start,limit
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields)
	p.SetFields(&fields)
	// FIXME do not set comp, get all tasks
	comp := int64(0)
	p.SetComp(&comp)
	num := limit
	p.SetNum(&num)
	start_ := &start
	p.SetStart(start_)
	if lastEditTime != nil {
		lastEditTime_ := int64(*lastEditTime)
		p.After = &lastEditTime_
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
	// XXX
	p := task.NewGetTasksDeletedPhpParams()
	if lastEditTime != nil {
		lastEditTime_ := int64(*lastEditTime)
		p.After = &lastEditTime_
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

// Create ...
func (s *taskService) Create(title string) (*models.Task, error) {
	t := models.Task{
		Title: title,
	}
	bytes, _ := json.Marshal([]models.Task{t})
	bytesS := (string)(bytes)
	p := task.NewPostTasksAddPhpParams()
	p.Tasks = &bytesS
	p.Fields = pointerutil.String(enums.TaskFields2String(enums.GeneralTaskFields))
	resp, err := s.cli.Task.PostTasksAddPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	// FIXME index
	return resp.Payload[0], err
}

// DeleteBatch ...
func (s *taskService) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
	p := task.NewPostTasksDeletePhpParams()
	idsString := funk.Map(ids, func(x int64) string {
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
	success := funk.Filter(resp.Payload, func(x *models.TaskDeleteItem) bool {
		return x.Ref == ""
	}).([]*models.TaskDeleteItem)
	successIds := funk.Map(success, func(x *models.TaskDeleteItem) int64 {
		return x.ID
	}).([]int64)

	failed := funk.Filter(resp.Payload, func(x *models.TaskDeleteItem) bool {
		return x.Ref != ""
	}).([]*models.TaskDeleteItem)
	return successIds, failed, nil

}

// Delete ...
func (s *taskService) Delete(id int64) error {
	t, err := s.FindById(id)
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

func (s *taskService) Edit(id int64, t *models.Task) (*models.Task, error) {
	t.ID = id
	bytes, _ := json.Marshal([]models.Task{*t})
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
	_, err := s.FindById(id)
	if err != nil {
		return nil, err
	}
	return s.Edit(id, &models.Task{
		Completed: time.Now().Unix(),
	})
}

func (s *taskService) UnComplete(id int64) (*models.Task, error) {
	_, err := s.FindById(id)
	if err != nil {
		return nil, err
	}
	return s.Edit(id, &models.Task{
		Completed: 0,
	})
}
