// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	broker "github.com/golangid/candi/broker"
	mock "github.com/stretchr/testify/mock"
)

// KafkaOptionFunc is an autogenerated mock type for the KafkaOptionFunc type
type KafkaOptionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *KafkaOptionFunc) Execute(_a0 *broker.KafkaBroker) {
	_m.Called(_a0)
}

// NewKafkaOptionFunc creates a new instance of KafkaOptionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKafkaOptionFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *KafkaOptionFunc {
	mock := &KafkaOptionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
