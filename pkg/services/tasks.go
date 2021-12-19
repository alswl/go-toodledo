package services

import (
	"encoding/json"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/task"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"strconv"
	"time"
)

// TaskService ...
type TaskService interface {
	FindById(id int64) (*models.Task, error)
	ListAll() ([]*models.Task, *models.PaginatedInfo, error)
	ListModifiedTimeIn(before, after time.Time, start, limit int, fields []enums.TaskField) ([]*models.Task, int, error)
	// TODO using opt
	Create(name string, options map[string]interface{}) (*models.Task, error)
	Delete(id int64) error
	DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error)
}

type taskService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

// NewTaskService ...
func NewTaskService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) TaskService {
	return &taskService{cli: cli, auth: auth}
}

// FindById ...
func (s *taskService) FindById(id int64) (*models.Task, error) {
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields)
	p.SetFields(&fields)
	p.SetID(&id)

	s.cli.Task.GetTasksGetPhp(p, s.auth)

	res, err := s.cli.Task.GetTasksGetPhp(p, s.auth)
	// XXX using multiple kind of payload item
	if err != nil || len(res.Payload) == 1 {
		return nil, err
	}
	//_ := res.Payload[0].(models.PaginatedInfo)
	bytes, _ := json.Marshal(res.Payload[1])
	var t models.Task
	json.Unmarshal(bytes, &t)
	return &t, nil
}

// ListAll ...
func (s *taskService) ListAll() ([]*models.Task, *models.PaginatedInfo, error) {
	// TODO using TaskQuery
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields)
	p.SetFields(&fields)
	comp := int64(0)
	p.SetComp(&comp)
	num := int64(10)
	p.SetNum(&num)

	s.cli.Task.GetTasksGetPhp(p, s.auth)

	res, err := s.cli.Task.GetTasksGetPhp(p, s.auth)
	if err != nil {
		return nil, nil, err
	}
	var paging models.PaginatedInfo
	bytes, _ := json.Marshal(res.Payload[0])
	json.Unmarshal(bytes, &paging)

	var tasks []*models.Task
	bytes, _ = json.Marshal(res.Payload[1:len(res.Payload)])
	json.Unmarshal(bytes, &tasks)
	return tasks, &paging, nil
}

// QueryModifiedTimeIn ...
func (s *taskService) ListModifiedTimeIn(before, after time.Time, start, limit int, fields []enums.TaskField) ([]*models.Task, int, error) {
	panic("implement me")
}

// Create ...
func (s *taskService) Create(title string, options map[string]interface{}) (*models.Task, error) {
	t := models.Task{
		Title: title,
	}
	// TODO options
	//for opt range options {
	//	opt(t)
	//}
	bytes, _ := json.Marshal([]models.Task{t})
	bytesS := (string)(bytes)
	p := task.NewPostTasksAddPhpParams()
	p.Tasks = &bytesS
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
		return fmt.Errorf("failed to delete task %d", id)
	}
	return nil
}
