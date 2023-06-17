package services_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alswl/go-toodledo/pkg/services"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	mockdal "github.com/alswl/go-toodledo/test/mock/pkg/dal"
	mockservices "github.com/alswl/go-toodledo/test/mock/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListAllByQueryNotDown(t *testing.T) {
	backend := mockdal.Backend{}
	var bytes [][]byte
	task := models.Task{
		Title: "task",
	}
	marshal, _ := json.Marshal(task)
	bytes = append(bytes, marshal)
	task = models.Task{
		Title:     "task-d-2000",
		Completed: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
	}
	marshal, _ = json.Marshal(task)
	bytes = append(bytes, marshal)

	backend.On("List", mock.Anything).Return(bytes, nil)
	s := services.NewTaskLocalExtService(&mockservices.TaskService{}, &mockservices.AccountExtService{}, &backend)

	query, err := s.ListAllByQuery(&queries.TaskListQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, query)
	assert.Len(t, query, 1)
}

func TestListAllByQueryDoneToday(t *testing.T) {
	backend := mockdal.Backend{}
	var bytes = [][]byte{}
	task := models.Task{
		Title: "task-not-done",
	}
	marshal, _ := json.Marshal(task)
	bytes = append(bytes, marshal)
	task = models.Task{
		Title:     "task-d-today",
		Completed: time.Now().Unix(),
	}
	marshal, _ = json.Marshal(task)
	bytes = append(bytes, marshal)

	backend.On("List", mock.Anything).Return(bytes, nil)
	s := services.NewTaskLocalExtService(&mockservices.TaskService{}, &mockservices.AccountExtService{}, &backend)

	query, err := s.ListAllByQuery(&queries.TaskListQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, query)
	assert.Len(t, query, 2)
}
