// Code generated by mockery v2.9.4. DO NOT EDIT.

package mockservices

import mock "github.com/stretchr/testify/mock"

// Cached is an autogenerated mock type for the Cached type
type Cached struct {
	mock.Mock
}

// LocalTruncate provides a mock function with given fields:
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
