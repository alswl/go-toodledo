// Code generated by mockery v2.12.2. DO NOT EDIT.

package mockservices

import (
	models "github.com/alswl/go-toodledo/pkg/models"
	mock "github.com/stretchr/testify/mock"

	queries "github.com/alswl/go-toodledo/pkg/models/queries"

	testing "testing"
)

// TaskCachedService is an autogenerated mock type for the TaskCachedService type
type TaskCachedService struct {
	mock.Mock
}

// Complete provides a mock function with given fields: id
func (_m *TaskCachedService) Complete(id int64) (*models.Task, error) {
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

// Create provides a mock function with given fields: name
func (_m *TaskCachedService) Create(name string) (*models.Task, error) {
	ret := _m.Called(name)

	var r0 *models.Task
	if rf, ok := ret.Get(0).(func(string) *models.Task); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateByQuery provides a mock function with given fields: query
func (_m *TaskCachedService) CreateByQuery(query *queries.TaskCreateQuery) (*models.Task, error) {
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
func (_m *TaskCachedService) Delete(id int64) error {
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
func (_m *TaskCachedService) DeleteBatch(ids []int64) ([]int64, []*models.TaskDeleteItem, error) {
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
func (_m *TaskCachedService) Edit(id int64, t *models.Task) (*models.Task, error) {
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
func (_m *TaskCachedService) FindById(id int64) (*models.Task, error) {
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

// List provides a mock function with given fields: start, limit
func (_m *TaskCachedService) List(start int64, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	ret := _m.Called(start, limit)

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func(int64, int64) []*models.Task); ok {
		r0 = rf(start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	var r1 *models.PaginatedInfo
	if rf, ok := ret.Get(1).(func(int64, int64) *models.PaginatedInfo); ok {
		r1 = rf(start, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.PaginatedInfo)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int64, int64) error); ok {
		r2 = rf(start, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAllByQuery provides a mock function with given fields: query
func (_m *TaskCachedService) ListAllByQuery(query *queries.TaskListQuery) ([]*models.Task, error) {
	ret := _m.Called(query)

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func(*queries.TaskListQuery) []*models.Task); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*queries.TaskListQuery) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeleted provides a mock function with given fields: lastEditTime
func (_m *TaskCachedService) ListDeleted(lastEditTime *int32) ([]*models.TaskDeleted, error) {
	ret := _m.Called(lastEditTime)

	var r0 []*models.TaskDeleted
	if rf, ok := ret.Get(0).(func(*int32) []*models.TaskDeleted); ok {
		r0 = rf(lastEditTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TaskDeleted)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*int32) error); ok {
		r1 = rf(lastEditTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWithChanged provides a mock function with given fields: lastEditTime, start, limit
func (_m *TaskCachedService) ListWithChanged(lastEditTime *int32, start int64, limit int64) ([]*models.Task, *models.PaginatedInfo, error) {
	ret := _m.Called(lastEditTime, start, limit)

	var r0 []*models.Task
	if rf, ok := ret.Get(0).(func(*int32, int64, int64) []*models.Task); ok {
		r0 = rf(lastEditTime, start, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	var r1 *models.PaginatedInfo
	if rf, ok := ret.Get(1).(func(*int32, int64, int64) *models.PaginatedInfo); ok {
		r1 = rf(lastEditTime, start, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.PaginatedInfo)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*int32, int64, int64) error); ok {
		r2 = rf(lastEditTime, start, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LocalClear provides a mock function with given fields:
func (_m *TaskCachedService) LocalClear() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PartialSync provides a mock function with given fields: lastEditTime
func (_m *TaskCachedService) PartialSync(lastEditTime *int32) error {
	ret := _m.Called(lastEditTime)

	var r0 error
	if rf, ok := ret.Get(0).(func(*int32) error); ok {
		r0 = rf(lastEditTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sync provides a mock function with given fields:
func (_m *TaskCachedService) Sync() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnComplete provides a mock function with given fields: id
func (_m *TaskCachedService) UnComplete(id int64) (*models.Task, error) {
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

// NewTaskCachedService creates a new instance of TaskCachedService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskCachedService(t testing.TB) *TaskCachedService {
	mock := &TaskCachedService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
