// Code generated by mockery v2.9.4. DO NOT EDIT.

package mockservices

import (
	enums "github.com/alswl/go-toodledo/pkg/models/enums"
	mock "github.com/stretchr/testify/mock"

	models "github.com/alswl/go-toodledo/pkg/models"

	queries "github.com/alswl/go-toodledo/pkg/models/queries"

	time "time"
)

// TaskService is an autogenerated mock type for the TaskService type
type TaskService struct {
	mock.Mock
}

// Complete provides a mock function with given fields: id
func (_m *TaskService) Complete(id int64) (*models.Task, error) {
	ret := _m.Called(id)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(int64) *models.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: name, options
func (_m *TaskService) Create(name string, options map[string]interface{}) (*models.Task, error) {
	ret := _m.Called(name, options)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) *models.Task); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateByQuery provides a mock function with given fields: query
func (_m *TaskService) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
	ret := _m.Called(query)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(*queries.TaskCreateQuery) *models.Task); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*queries.TaskCreateQuery) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *TaskService) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBatch provides a mock function with given fields: ids
func (_m *TaskService) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
	ret := _m.Called(ids)

	var r0 []int64
	if rf, ok := ret.Get(0).(func([]int64) []int64); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 []*models.TaskDeleteItem
	if rf, ok := ret.Get(1).(func([]int64) []*models.TaskDeleteItem); ok {
		r1 = rf(ids)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*models.TaskDeleteItem)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]int64) error); ok {
		r2 = rf(ids)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Edit provides a mock function with given fields: id, t
func (_m *TaskService) Edit(id int64, t *models.Task) (*models.Task, error) {
	ret := _m.Called(id, t)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(int64, *models.Task) *models.Task); ok {
		r0 = rf(id, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *models.Task) error); ok {
		r1 = rf(id, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *TaskService) FindById(id int64) (*models.Task, error) {
	ret := _m.Called(id)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(int64) *models.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAll provides a mock function with given fields:
func (_m *TaskService) ListAll() ([]*models.Task, *models.PaginatedInfo, error) {
	ret := _m.Called()

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func() []*models.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	var r1 *models.PaginatedInfo
	if rf, ok := ret.Get(1).(func() *models.PaginatedInfo); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.PaginatedInfo)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListByQuery provides a mock function with given fields: query
func (_m *TaskService) ListByQuery(query *queries.TaskSearchQuery) ([]*models.Task, *models.PaginatedInfo, error) {
	ret := _m.Called(query)

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func(*queries.TaskSearchQuery) []*models.Task); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	var r1 *models.PaginatedInfo
	if rf, ok := ret.Get(1).(func(*queries.TaskSearchQuery) *models.PaginatedInfo); ok {
		r1 = rf(query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.PaginatedInfo)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*queries.TaskSearchQuery) error); ok {
		r2 = rf(query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListModifiedTimeIn provides a mock function with given fields: before, after, start, limit, fields
func (_m *TaskService) ListModifiedTimeIn(before time.Time, after time.Time, start int, limit int, fields []enums.TaskField) ([]*models.Task, int, error) {
	ret := _m.Called(before, after, start, limit, fields)

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, int, []enums.TaskField) []*models.Task); ok {
		r0 = rf(before, after, start, limit, fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(time.Time, time.Time, int, int, []enums.TaskField) int); ok {
		r1 = rf(before, after, start, limit, fields)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(time.Time, time.Time, int, int, []enums.TaskField) error); ok {
		r2 = rf(before, after, start, limit, fields)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UnComplete provides a mock function with given fields: id
func (_m *TaskService) UnComplete(id int64) (*models.Task, error) {
	ret := _m.Called(id)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(int64) *models.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
