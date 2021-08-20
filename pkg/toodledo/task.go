package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"time"
)

// TODO
type TaskService interface {
	GetById(ctx context.Context, id int) (models.Task, Response, error)
	Query(ctx context.Context, query models.TaskQuery) ([]models.Task, Response, error)
	GetDeleted(ctx context.Context) ([]models.Task, Response, error)
	Add(ctx context.Context, taskAdds []models.TaskAdd) (models.Task, Response, error)
	Edit(ctx context.Context, id int, name string, private bool) (models.Task, Response, error)
	Delete(ctx context.Context, after time.Time) (Response, error)
}

type taskService Service

func (t *taskService) GetById(ctx context.Context, id int) (models.Task, Response, error) {
	panic("implement me")
}

func (t *taskService) Query(ctx context.Context, query models.TaskQuery) ([]models.Task, Response, error) {
	panic("implement me")
}

func (t *taskService) GetDeleted(ctx context.Context) ([]models.Task, Response, error) {
	panic("implement me")
}

func (t *taskService) Add(ctx context.Context, taskAdds []models.TaskAdd) (models.Task, Response, error) {
	panic("implement me")
}

func (t *taskService) Edit(ctx context.Context, id int, name string, private bool) (models.Task, Response, error) {
	panic("implement me")
}

func (t *taskService) Delete(ctx context.Context, after time.Time) (Response, error) {
	panic("implement me")
}
