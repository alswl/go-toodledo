package services

import "github.com/alswl/go-toodledo/pkg/models"

type TaskRichService interface {
	FindByIdRich(id int64) (*models.RichTask, error)
	RichThem(tasks []*models.Task) ([]*models.RichTask, error)
}

type taskRichService struct {
	taskSvc    TaskCachedService
	folderSvc  FolderCachedService
	contextSvc ContextCachedService
	goalSvc    GoalCachedService
}

func NewTaskRichService(taskSvc TaskCachedService, folderSvc FolderCachedService, contextSvc ContextCachedService, goalSvc GoalCachedService) TaskRichService {
	return &taskRichService{taskSvc: taskSvc, folderSvc: folderSvc, contextSvc: contextSvc, goalSvc: goalSvc}
}

func (s *taskRichService) FindByIdRich(id int64) (*models.RichTask, error) {
	t, err := s.taskSvc.FindById(id)
	if err != nil {
		return nil, err
	}
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
	for _, task := range tasks {
		rt, err := s.FindByIdRich(task.ID)
		if err != nil {
			return nil, err
		}
		rts = append(rts, rt)
	}
	return rts, nil
}
