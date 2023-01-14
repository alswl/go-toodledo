//go:build integration
// +build integration

package services

import (
	"github.com/alswl/go-toodledo/test/suites/itinjector"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTaskServiceFindById(t *testing.T) {
	app, err := itinjector.InitTUIApp()
	assert.NotNil(t, err)
	svc := app.TaskSvc
	assert.NotNil(t, err)

	task, err := svc.FindByID(273321713)
	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func TestTaskListDeleted(t *testing.T) {
	app, err := itinjector.InitTUIApp()
	assert.NoError(t, err)
	svc := app.TaskSvc

	time := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	unix := int32(time.Unix())
	tasks, err := svc.ListDeleted(&unix)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
}
