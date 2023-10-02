// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StructValidator is an autogenerated mock type for the StructValidator type
type StructValidator struct {
	mock.Mock
}

// ValidateStruct provides a mock function with given fields: data
func (_m *StructValidator) ValidateStruct(data interface{}) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStructValidator creates a new instance of StructValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStructValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *StructValidator {
	mock := &StructValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
