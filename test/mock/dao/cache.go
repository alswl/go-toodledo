// Code generated by mockery v2.9.4. DO NOT EDIT.

package mockdao

import mock "github.com/stretchr/testify/mock"

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// Find provides a mock function with given fields: identity
func (_m *Cache) Find(identity string) (interface{}, error) {
	ret := _m.Called(identity)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(identity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Invalid provides a mock function with given fields:
func (_m *Cache) Invalid() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsExpired provides a mock function with given fields:
func (_m *Cache) IsExpired() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListAll provides a mock function with given fields: _a0
func (_m *Cache) ListAll(_a0 []interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
