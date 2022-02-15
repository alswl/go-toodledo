// Code generated by mockery v2.10.0. DO NOT EDIT.

package mockservices

import mock "github.com/stretchr/testify/mock"

// Cached is an autogenerated mock type for the Cached type
type Cached struct {
	mock.Mock
}

// LocalClear provides a mock function with given fields:
func (_m *Cached) LocalClear() error {
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
func (_m *Cached) PartialSync(lastEditTime *int32) error {
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
func (_m *Cached) Sync() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
