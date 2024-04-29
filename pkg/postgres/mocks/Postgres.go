// Code generated by mockery v2.42.2. DO NOT EDIT.

package postgresmocks

import (
	context "context"

	pgconn "github.com/jackc/pgconn"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	postgres "github.com/kenyako/platform_common/pkg/postgres"
)

// Postgres is an autogenerated mock type for the Postgres type
type Postgres struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx, txOptions
func (_m *Postgres) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (postgres.Tx, error) {
	ret := _m.Called(ctx, txOptions)

	if len(ret) == 0 {
		panic("no return value specified for BeginTx")
	}

	var r0 postgres.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) (postgres.Tx, error)); ok {
		return rf(ctx, txOptions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.TxOptions) postgres.Tx); ok {
		r0 = rf(ctx, txOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(postgres.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.TxOptions) error); ok {
		r1 = rf(ctx, txOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Postgres) Close() {
	_m.Called()
}

// ExecContext provides a mock function with given fields: ctx, q, args
func (_m *Postgres) ExecContext(ctx context.Context, q postgres.Query, args ...interface{}) (pgconn.CommandTag, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, postgres.Query, ...interface{}) (pgconn.CommandTag, error)); ok {
		return rf(ctx, q, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, postgres.Query, ...interface{}) pgconn.CommandTag); ok {
		r0 = rf(ctx, q, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgconn.CommandTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, postgres.Query, ...interface{}) error); ok {
		r1 = rf(ctx, q, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *Postgres) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryContext provides a mock function with given fields: ctx, q, args
func (_m *Postgres) QueryContext(ctx context.Context, q postgres.Query, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryContext")
	}

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, postgres.Query, ...interface{}) (pgx.Rows, error)); ok {
		return rf(ctx, q, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, postgres.Query, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, q, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, postgres.Query, ...interface{}) error); ok {
		r1 = rf(ctx, q, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryRowContext provides a mock function with given fields: ctx, q, args
func (_m *Postgres) QueryRowContext(ctx context.Context, q postgres.Query, args ...interface{}) pgx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRowContext")
	}

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func(context.Context, postgres.Query, ...interface{}) pgx.Row); ok {
		r0 = rf(ctx, q, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// ScanAllContext provides a mock function with given fields: ctx, dest, q, args
func (_m *Postgres) ScanAllContext(ctx context.Context, dest interface{}, q postgres.Query, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ScanAllContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, postgres.Query, ...interface{}) error); ok {
		r0 = rf(ctx, dest, q, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScanOneContext provides a mock function with given fields: ctx, dest, q, args
func (_m *Postgres) ScanOneContext(ctx context.Context, dest interface{}, q postgres.Query, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ScanOneContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, postgres.Query, ...interface{}) error); ok {
		r0 = rf(ctx, dest, q, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPostgres creates a new instance of Postgres. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostgres(t interface {
	mock.TestingT
	Cleanup(func())
}) *Postgres {
	mock := &Postgres{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
