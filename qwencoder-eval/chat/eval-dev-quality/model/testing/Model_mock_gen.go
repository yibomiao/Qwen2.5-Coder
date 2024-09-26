// Code generated by mockery v2.40.3. DO NOT EDIT.

package modeltesting

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/symflower/eval-dev-quality/model"
)

// MockModel is an autogenerated mock type for the Model type
type MockModel struct {
	mock.Mock
}

// ID provides a mock function with given fields:
func (_m *MockModel) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MetaInformation provides a mock function with given fields:
func (_m *MockModel) MetaInformation() *model.MetaInformation {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MetaInformation")
	}

	var r0 *model.MetaInformation
	if rf, ok := ret.Get(0).(func() *model.MetaInformation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MetaInformation)
		}
	}

	return r0
}

// NewMockModel creates a new instance of MockModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockModel {
	mock := &MockModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
