package services

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
)

type TaskRichService interface {
	Find(id int64) (*models.RichTask, error)
	Rich(tasks *models.Task) (*models.RichTask, error)
	RichThem(tasks []*models.Task) ([]*models.RichTask, error)
}

type TaskRichCachedService = TaskRichService

type taskRichService struct {
	taskSvc    TaskService
	folderSvc  FolderService
	contextSvc ContextService
	goalSvc    GoalService
	logger     logrus.FieldLogger
}

// NewTaskRichService create rich service with cached service(except task service)
func NewTaskRichService(taskSvc TaskService, folderSvc FolderCachedService, contextSvc ContextCachedService, goalSvc GoalCachedService, logger logrus.FieldLogger) TaskRichCachedService {
	return &taskRichService{taskSvc: taskSvc, folderSvc: folderSvc, contextSvc: contextSvc, goalSvc: goalSvc, logger: logger}
}

func NewTaskRichCachedService(taskSvc TaskCachedService, folderSvc FolderCachedService, contextSvc ContextCachedService, goalSvc GoalCachedService, logger logrus.FieldLogger) TaskRichCachedService {
	return &taskRichService{taskSvc: taskSvc, folderSvc: folderSvc, contextSvc: contextSvc, goalSvc: goalSvc, logger: logger}
}

func (s *taskRichService) Find(id int64) (*models.RichTask, error) {
	// FIXME deprecated, using Rich()
	t, err := s.taskSvc.FindById(id)
	if err != nil {
		return nil, err
	}
	return s.Rich(t)
}

func (s *taskRichService) Rich(t *models.Task) (*models.RichTask, error) {
	var context = &models.Context{}
	if t.Context != 0 {
		context, _ = s.contextSvc.FindByID(t.Context)
	}
	var folder = &models.Folder{}
	if t.Folder != 0 {
		folder, _ = s.folderSvc.FindByID(t.Folder)
	}
	var goal = &models.Goal{}
	if t.Goal != 0 {
		goal, _ = s.goalSvc.FindByID(t.Goal)
	}

	rt := &models.RichTask{
		Task:       *t,
		TheContext: *context,
		TheFolder:  *folder,
		TheGoal:    *goal,
	}
	return rt, nil
}

func (s *taskRichService) RichThem(tasks []*models.Task) ([]*models.RichTask, error) {
	var rts []*models.RichTask
	// FIXME rich with context, folder, goal
	for _, task := range tasks {
		rt, err := s.Find(task.ID)
		if err != nil {
			return nil, err
		}
		rts = append(rts, rt)
	}
	return rts, nil
}
