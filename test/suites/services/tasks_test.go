//go:build integration
// +build integration

package services

import (
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/test/suites"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskServiceFindById(t *testing.T) {
	auth := suites.AuthForTest()

	svc := services.NewTaskService(auth)
	task, err := svc.FindById(273321713)
	assert.NoError(t, err)
	assert.NotNil(t, task)
}
