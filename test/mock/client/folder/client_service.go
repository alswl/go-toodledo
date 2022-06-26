// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockfolder

import (
	folder "github.com/alswl/go-toodledo/pkg/client/folder"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetFoldersGetPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetFoldersGetPhp(params *folder.GetFoldersGetPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...folder.ClientOption) (*folder.GetFoldersGetPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *folder.GetFoldersGetPhpOK
	if rf, ok := ret.Get(0).(func(*folder.GetFoldersGetPhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) *folder.GetFoldersGetPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*folder.GetFoldersGetPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*folder.GetFoldersGetPhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostFoldersAddPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostFoldersAddPhp(params *folder.PostFoldersAddPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...folder.ClientOption) (*folder.PostFoldersAddPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *folder.PostFoldersAddPhpOK
	if rf, ok := ret.Get(0).(func(*folder.PostFoldersAddPhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) *folder.PostFoldersAddPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*folder.PostFoldersAddPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*folder.PostFoldersAddPhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostFoldersDeletePhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostFoldersDeletePhp(params *folder.PostFoldersDeletePhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...folder.ClientOption) (*folder.PostFoldersDeletePhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *folder.PostFoldersDeletePhpOK
	if rf, ok := ret.Get(0).(func(*folder.PostFoldersDeletePhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) *folder.PostFoldersDeletePhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*folder.PostFoldersDeletePhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*folder.PostFoldersDeletePhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostFoldersEditPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostFoldersEditPhp(params *folder.PostFoldersEditPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...folder.ClientOption) (*folder.PostFoldersEditPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *folder.PostFoldersEditPhpOK
	if rf, ok := ret.Get(0).(func(*folder.PostFoldersEditPhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) *folder.PostFoldersEditPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*folder.PostFoldersEditPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*folder.PostFoldersEditPhpParams, runtime.ClientAuthInfoWriter, ...folder.ClientOption) error); ok {
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
