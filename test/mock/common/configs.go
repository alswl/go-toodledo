// Code generated by mockery v2.9.4. DO NOT EDIT.

package mockcommon

import (
	common "github.com/alswl/go-toodledo/pkg/common"
	mock "github.com/stretchr/testify/mock"
)

// Configs is an autogenerated mock type for the Configs type
type Configs struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *Configs) Get() *common.ToodledoConfig {
	ret := _m.Called()

	var r0 *common.ToodledoConfig
	if rf, ok := ret.Get(0).(func() *common.ToodledoConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.ToodledoConfig)
		}
	}

	return r0
}
