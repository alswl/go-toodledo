package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/task"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"time"
)

type TaskService interface {
	FindById(id int64) (*models.Task, error)

	QueryAll() ([]*models.Task, *models.PaginatedInfo, error)

	QueryModifiedTimeIn(before, after time.Time, start, limit int, fields []enums.TaskField) ([]*models.Task, int, error)
}

type taskService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

func NewTaskService(auth runtime.ClientAuthInfoWriter) *taskService {
	return &taskService{
		cli:  client.NewHTTPClient(strfmt.NewFormats()),
		auth: auth,
	}
}

func ProvideTaskService(auth runtime.ClientAuthInfoWriter) TaskService {
	return NewTaskService(auth)
}

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

func (s *taskService) QueryModifiedTimeIn(before, after time.Time, start, limit int, fields []enums.TaskField) ([]*models.Task, int, error) {
	panic("implement me")
}
