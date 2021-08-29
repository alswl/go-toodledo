// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocktoodledo

import (
	context "context"

	models "github.com/alswl/go-toodledo/pkg/toodledo/models"
	mock "github.com/stretchr/testify/mock"

	toodledo "github.com/alswl/go-toodledo/pkg/toodledo"
)

// SavedSearchService is an autogenerated mock type for the SavedSearchService type
type SavedSearchService struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx
func (_m *SavedSearchService) Get(ctx context.Context) (models.SavedSearch, toodledo.Response, error) {
	ret := _m.Called(ctx)

	var r0 models.SavedSearch
	if rf, ok := ret.Get(0).(func(context.Context) models.SavedSearch); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(models.SavedSearch)
	}

	var r1 toodledo.Response
	if rf, ok := ret.Get(1).(func(context.Context) toodledo.Response); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(toodledo.Response)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
