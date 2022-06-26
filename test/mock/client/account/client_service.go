// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockaccount

import (
	account "github.com/alswl/go-toodledo/pkg/client/account"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetAccountGetPhp provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetAccountGetPhp(params *account.GetAccountGetPhpParams, authInfo runtime.ClientAuthInfoWriter, opts ...account.ClientOption) (*account.GetAccountGetPhpOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *account.GetAccountGetPhpOK
	if rf, ok := ret.Get(0).(func(*account.GetAccountGetPhpParams, runtime.ClientAuthInfoWriter, ...account.ClientOption) *account.GetAccountGetPhpOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.GetAccountGetPhpOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*account.GetAccountGetPhpParams, runtime.ClientAuthInfoWriter, ...account.ClientOption) error); ok {
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
