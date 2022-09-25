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
	_, err := itinjector.InitTUIApp()
	assert.NotNil(t, err)
	svc, err := itinjector.InitTaskService()
	assert.NotNil(t, err)

	task, err := svc.FindById(273321713)
	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func TestTaskListDeleted(t *testing.T) {
	svc, err := itinjector.InitTaskService()
	assert.NoError(t, err)

	time := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	unix := int32(time.Unix())
	tasks, err := svc.ListDeleted(&unix)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
}
