// Code generated by mockery v2.9.4. DO NOT EDIT.

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

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
