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
	logger     logrus.FieldLogger
	taskSvc    TaskService
	folderSvc  FolderService
	contextSvc ContextService
	goalSvc    GoalPersistenceService
}

// NewTaskRichService returns a new TaskRichService.
// it will not cache the attribute when rich the task, and query the attribute from remote
// Please use the NewTaskRichPersistenceService first.
func NewTaskRichService(
	taskSvc TaskService,
	folderSvc FolderService,
	contextSvc ContextService,
	goalSvc GoalPersistenceService,
	logger logrus.FieldLogger,
) TaskRichService {
	return &taskRichService{
		taskSvc:    taskSvc,
		folderSvc:  folderSvc,
		contextSvc: contextSvc,
		goalSvc:    goalSvc,
		logger:     logger,
	}
}

// NewTaskRichPersistenceService returns a new TaskRichCachedService.
// svc inside will cache the attribute when rich the task.
func NewTaskRichPersistenceService(
	taskSvc TaskPersistenceExtService,
	folderSvc FolderPersistenceService,
	contextSvc ContextPersistenceService,
	goalSvc GoalPersistenceService,
	logger logrus.FieldLogger,
) TaskRichCachedService {
	return &taskRichService{
		taskSvc:    taskSvc,
		folderSvc:  folderSvc,
		contextSvc: contextSvc,
		goalSvc:    goalSvc,
		logger:     logger,
	}
}

func (s *taskRichService) Find(id int64) (*models.RichTask, error) {
	// FIXME deprecated, using Rich()
	t, err := s.taskSvc.FindByID(id)
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
		TheContext: context,
		TheFolder:  folder,
		TheGoal:    goal,
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
