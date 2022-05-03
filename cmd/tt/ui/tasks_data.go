package ui

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/bxcodec/faker/v3"
	"github.com/thoas/go-funk"
)

func AllTasksMock() ([]*models.RichTask, error) {
	tasks := make([]models.RichTask, 50)
	faker.FakeData(&tasks)

	ts := funk.Map(tasks, func(x models.RichTask) *models.RichTask {
		return &x
	}).([]*models.RichTask)

	return ts, nil
}
