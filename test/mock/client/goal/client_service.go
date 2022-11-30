// Code generated by mockery v2.15.0. DO NOT EDIT.

package mockgoal

import (
	goal "github.com/alswl/go-toodledo/pkg/client/goal"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetGoalsGetPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetGoalsGetPhp(params *goal.GetGoalsGetPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...goal.ClientOption) (*goal.GetGoalsGetPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *goal.GetGoalsGetPhpOK
	if rf, ok := ret.Get(0).(func(*goal.GetGoalsGetPhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) *goal.GetGoalsGetPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goal.GetGoalsGetPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*goal.GetGoalsGetPhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostGoalsAddPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostGoalsAddPhp(params *goal.PostGoalsAddPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...goal.ClientOption) (*goal.PostGoalsAddPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *goal.PostGoalsAddPhpOK
	if rf, ok := ret.Get(0).(func(*goal.PostGoalsAddPhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) *goal.PostGoalsAddPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goal.PostGoalsAddPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*goal.PostGoalsAddPhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostGoalsDeletePhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostGoalsDeletePhp(params *goal.PostGoalsDeletePhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...goal.ClientOption) (*goal.PostGoalsDeletePhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *goal.PostGoalsDeletePhpOK
	if rf, ok := ret.Get(0).(func(*goal.PostGoalsDeletePhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) *goal.PostGoalsDeletePhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goal.PostGoalsDeletePhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*goal.PostGoalsDeletePhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostGoalsEditPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostGoalsEditPhp(params *goal.PostGoalsEditPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...goal.ClientOption) (*goal.PostGoalsEditPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *goal.PostGoalsEditPhpOK
	if rf, ok := ret.Get(0).(func(*goal.PostGoalsEditPhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) *goal.PostGoalsEditPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goal.PostGoalsEditPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*goal.PostGoalsEditPhpParams, runtime.ClientAuthInfoWriter, ...goal.ClientOption) error); ok {
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
