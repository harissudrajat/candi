// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	taskqueueworker "github.com/golangid/candi/codebase/app/task_queue_worker"
	mock "github.com/stretchr/testify/mock"
)

// QueueStorage is an autogenerated mock type for the QueueStorage type
type QueueStorage struct {
	mock.Mock
}

// Clear provides a mock function with given fields: ctx, taskName
func (_m *QueueStorage) Clear(ctx context.Context, taskName string) {
	_m.Called(ctx, taskName)
}

// NextJob provides a mock function with given fields: ctx, taskName
func (_m *QueueStorage) NextJob(ctx context.Context, taskName string) string {
	ret := _m.Called(ctx, taskName)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, taskName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *QueueStorage) Ping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PopJob provides a mock function with given fields: ctx, taskName
func (_m *QueueStorage) PopJob(ctx context.Context, taskName string) string {
	ret := _m.Called(ctx, taskName)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, taskName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PushJob provides a mock function with given fields: ctx, job
func (_m *QueueStorage) PushJob(ctx context.Context, job *taskqueueworker.Job) int64 {
	ret := _m.Called(ctx, job)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *taskqueueworker.Job) int64); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *QueueStorage) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewQueueStorage creates a new instance of QueueStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueueStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueueStorage {
	mock := &QueueStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
