package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/task"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums"
	"github.com/go-openapi/runtime"
	"time"
)

// TaskService ...
type TaskService interface {
	FindById(id int64) (*models.Task, error)
	QueryAll() ([]*models.Task, *models.PaginatedInfo, error)
	QueryModifiedTimeIn(before, after time.Time, start, limit int, fields []enums.TaskField) ([]*models.Task, int, error)
	Create(name string, options map[string]interface{}) (*models.Task, error)
}

type taskService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

// NewTaskService ...
func NewTaskService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) TaskService {
	return &taskService{
		cli:  cli,
		auth: auth,
	}
}

// FindById ...
func (s *taskService) FindById(id int64) (*models.Task, error) {
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields)
	p.SetFields(&fields)
	p.SetID(&id)

	s.cli.Task.GetTasksGetPhp(p, s.auth)

	res, err := s.cli.Task.GetTasksGetPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	//_ := res.Payload[0].(models.PaginatedInfo)
	bytes, _ := json.Marshal(res.Payload[1])
	var t models.Task
	json.Unmarshal(bytes, &t)
	return &t, nil
}

// QueryAll ...
func (s *taskService) QueryAll() ([]*models.Task, *models.PaginatedInfo, error) {
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
func (s *taskService) QueryModifiedTimeIn(before, after time.Time, start, limit int, fields []enums.TaskField) ([]*models.Task, int, error) {
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
