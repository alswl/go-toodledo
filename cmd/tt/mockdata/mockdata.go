package mockdata

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/bxcodec/faker/v3"
	"github.com/thoas/go-funk"
)

func AllTasksMock() ([]*models.RichTask, error) {
	tasks := make([]models.RichTask, 500)
	_ = faker.FakeData(&tasks)

	ts := funk.Map(tasks, func(x models.RichTask) *models.RichTask {
		return &x
	}).([]*models.RichTask)

	return ts, nil
}

func ListContexts() []models.Context {
	contexts := make([]models.Context, 10)
	_ = faker.FakeData(&contexts)
	return contexts
}
