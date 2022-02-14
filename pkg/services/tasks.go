package services

import (
	"encoding/json"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/task"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"strconv"
	"time"
)

var DefaultFieldsInResponse = "folder,star,context,tag,goal,repeat,startdate,starttime,duedate,duetime,priority,length"

// TaskService ...
type TaskService interface {
	FindById(id int64) (*models.Task, error)
	List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error)
	Create(name string) (*models.Task, error)
	CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error)
	Delete(id int64) error
	DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error)
	Edit(id int64, t *models.Task) (*models.Task, error)
	Complete(id int64) (*models.Task, error)
	UnComplete(id int64) (*models.Task, error)
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
	// TODO using multiple kind of payload item
	if err != nil || len(res.Payload) == 1 {
		return nil, err
	}
	//_ := res.Payload[0].(models.PaginatedInfo)
	bytes, _ := json.Marshal(res.Payload[1])
	var t models.Task
	json.Unmarshal(bytes, &t)
	return &t, nil
}

// listAllRemote ...
func (s *taskService) List(start, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	// TODO using TaskQuery,before, after,start,limit
	p := task.NewGetTasksGetPhpParams()
	fields := enums.TaskFields2String(enums.GeneralTaskFields)
	p.SetFields(&fields)
	comp := int64(0)
	p.SetComp(&comp)
	num := limit
	p.SetNum(&num)
	start_ := &start
	p.SetStart(start_)

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

func (s *taskService) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
	ts := []models.Task{*query.ToModel()}
	bytes, _ := json.Marshal(ts)
	bytesS := (string)(bytes)
	p := task.NewPostTasksAddPhpParams()
	p.Tasks = &bytesS
	p.Fields = &DefaultFieldsInResponse
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

// Edit ...
func (s *taskService) Edit(id int64, t *models.Task) (*models.Task, error) {
	bytes, _ := json.Marshal([]models.Task{*t})
	bytesS := (string)(bytes)
	p := task.NewPostTasksEditPhpParams()
	p.Tasks = &bytesS
	resp, err := s.cli.Task.PostTasksEditPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	// FIXME index
	return resp.Payload[0], err
}

func (s *taskService) Complete(id int64) (*models.Task, error) {
	t, err := s.FindById(id)
	if err != nil {
		return nil, err
	}
	return s.Edit(id, &models.Task{
		ID:        t.ID,
		Completed: time.Now().Unix(),
	})
}

func (s *taskService) UnComplete(id int64) (*models.Task, error) {
	t, err := s.FindById(id)
	if err != nil {
		return nil, err
	}
	return s.Edit(id, &models.Task{
		ID:        t.ID,
		Completed: 0,
	})
}

//if b.folder != "" && query.FolderID == 0 {
//	folder, err := b.folderSvc.Find(b.folder)
//	if err != nil {
//		return errors.Wrap(err, "failed to find folder")
//	}
//	query.FolderID = folder.ID
//}
//
//if b.context != "" && query.ContextID == 0 {
//	context, err := b.contextSvc.Find(b.context)
//	if err != nil {
//		return errors.Wrap(err, "failed to find context")
//	}
//	query.ContextID = context.ID
//}
//if b.goal != "" && query.GoalID == 0 {
//	goal, err := b.goalSvc.Find(b.goal)
//	if err != nil {
//		return errors.Wrap(err, "failed to find goal")
//	}
//	query.GoalID = goal.ID
//}
//
//return nil
