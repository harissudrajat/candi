// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// SQLDatabase is an autogenerated mock type for the SQLDatabase type
type SQLDatabase struct {
	mock.Mock
}

// Disconnect provides a mock function with given fields: ctx
func (_m *SQLDatabase) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Health provides a mock function with given fields:
func (_m *SQLDatabase) Health() map[string]error {
	ret := _m.Called()

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// ReadDB provides a mock function with given fields:
func (_m *SQLDatabase) ReadDB() *sql.DB {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// WriteDB provides a mock function with given fields:
func (_m *SQLDatabase) WriteDB() *sql.DB {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// NewSQLDatabase creates a new instance of SQLDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSQLDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *SQLDatabase {
	mock := &SQLDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
