// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocktask

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	task "github.com/alswl/go-toodledo/pkg/client/task"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetTasksDeletedPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetTasksDeletedPhp(params *task.GetTasksDeletedPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...task.ClientOption) (*task.GetTasksDeletedPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *task.GetTasksDeletedPhpOK
	if rf, ok := ret.Get(0).(func(*task.GetTasksDeletedPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) *task.GetTasksDeletedPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.GetTasksDeletedPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*task.GetTasksDeletedPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasksGetPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetTasksGetPhp(params *task.GetTasksGetPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...task.ClientOption) (*task.GetTasksGetPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *task.GetTasksGetPhpOK
	if rf, ok := ret.Get(0).(func(*task.GetTasksGetPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) *task.GetTasksGetPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.GetTasksGetPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*task.GetTasksGetPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostTasksAddPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostTasksAddPhp(params *task.PostTasksAddPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...task.ClientOption) (*task.PostTasksAddPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *task.PostTasksAddPhpOK
	if rf, ok := ret.Get(0).(func(*task.PostTasksAddPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) *task.PostTasksAddPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.PostTasksAddPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*task.PostTasksAddPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostTasksDeletePhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostTasksDeletePhp(params *task.PostTasksDeletePhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...task.ClientOption) (*task.PostTasksDeletePhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *task.PostTasksDeletePhpOK
	if rf, ok := ret.Get(0).(func(*task.PostTasksDeletePhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) *task.PostTasksDeletePhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.PostTasksDeletePhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*task.PostTasksDeletePhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostTasksEditPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostTasksEditPhp(params *task.PostTasksEditPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...task.ClientOption) (*task.PostTasksEditPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *task.PostTasksEditPhpOK
	if rf, ok := ret.Get(0).(func(*task.PostTasksEditPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) *task.PostTasksEditPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.PostTasksEditPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*task.PostTasksEditPhpParams, runtime.ClientAuthInfoWriter, ...task.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientService(t mockConstructorTestingTNewClientService) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
