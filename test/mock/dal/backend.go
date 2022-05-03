// Code generated by mockery v2.12.1. DO NOT EDIT.

package mockdal

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Backend) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: bucket, key
func (_m *Backend) Get(bucket string, key string) ([]byte, error) {
	ret := _m.Called(bucket, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(bucket, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bucket, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Keys provides a mock function with given fields: bucket
func (_m *Backend) Keys(bucket string) ([]string, error) {
	ret := _m.Called(bucket)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: bucket
func (_m *Backend) List(bucket string) ([][]byte, error) {
	ret := _m.Called(bucket)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(string) [][]byte); ok {
		r0 = rf(bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Path provides a mock function with given fields: key
func (_m *Backend) Path(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Put provides a mock function with given fields: bucket, key, value
func (_m *Backend) Put(bucket string, key string, value []byte) error {
	ret := _m.Called(bucket, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []byte) error); ok {
		r0 = rf(bucket, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: bucket, key
func (_m *Backend) Remove(bucket string, key string) error {
	ret := _m.Called(bucket, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(bucket, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Truncate provides a mock function with given fields: bucket
func (_m *Backend) Truncate(bucket string) error {
	ret := _m.Called(bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBackend creates a new instance of Backend. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewBackend(t testing.TB) *Backend {
	mock := &Backend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
