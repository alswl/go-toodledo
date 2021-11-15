//go:build integration
// +build integration

package services

import (
	"github.com/alswl/go-toodledo/test/suites/itinjector"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskServiceFindById(t *testing.T) {
	app, err := itinjector.InitApp()

	svc := app.TaskSvc
	task, err := svc.FindById(273321713)
	assert.NoError(t, err)
	assert.NotNil(t, task)
}
