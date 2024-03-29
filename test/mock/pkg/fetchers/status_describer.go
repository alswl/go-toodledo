// Code generated by mockery v2.38.0. DO NOT EDIT.

package mockfetchers

import mock "github.com/stretchr/testify/mock"

// StatusDescriber is an autogenerated mock type for the StatusDescriber type
type StatusDescriber struct {
	mock.Mock
}

// Error provides a mock function with given fields: err
func (_m *StatusDescriber) Error(err error) {
	_m.Called(err)
}

// OnError provides a mock function with given fields: _a0
func (_m *StatusDescriber) OnError(_a0 func(error) error) {
	_m.Called(_a0)
}

// OnSuccess provides a mock function with given fields: _a0
func (_m *StatusDescriber) OnSuccess(_a0 func() error) {
	_m.Called(_a0)
}

// OnSyncing provides a mock function with given fields: _a0
func (_m *StatusDescriber) OnSyncing(_a0 func() error) {
	_m.Called(_a0)
}

// Progress provides a mock function with given fields:
func (_m *StatusDescriber) Progress() (int, int) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Progress")
	}

	var r0 int
	var r1 int
	if rf, ok := ret.Get(0).(func() (int, int)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// SetProgress provides a mock function with given fields: current, total
func (_m *StatusDescriber) SetProgress(current int, total int) {
	_m.Called(current, total)
}

// Success provides a mock function with given fields:
func (_m *StatusDescriber) Success() {
	_m.Called()
}

// Syncing provides a mock function with given fields:
func (_m *StatusDescriber) Syncing() {
	_m.Called()
}

// NewStatusDescriber creates a new instance of StatusDescriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatusDescriber(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatusDescriber {
	mock := &StatusDescriber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
