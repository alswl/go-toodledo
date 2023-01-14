package mockdata

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/bxcodec/faker/v3"
	"github.com/thoas/go-funk"
)

func AllTasksMock() ([]*models.RichTask, error) {
	// TODO no works now
	const defaultCount = 500
	tasks := make([]models.RichTask, defaultCount)
	_ = faker.FakeData(&tasks)

	ts, _ := funk.Map(tasks, func(x models.RichTask) *models.RichTask {
		return &x
	}).([]*models.RichTask)

	return ts, nil
}

func ListContexts() []models.Context {
	const size = 10
	contexts := make([]models.Context, size)
	_ = faker.FakeData(&contexts)
	return contexts
}
