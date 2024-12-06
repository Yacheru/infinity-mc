// Code generated by mockery v2.49.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AuthMongoRepository is an autogenerated mock type for the AuthMongoRepository type
type AuthMongoRepository struct {
	mock.Mock
}

// GetCode provides a mock function with given fields: ctx, email, code
func (_m *AuthMongoRepository) GetCode(ctx context.Context, email string, code int) error {
	ret := _m.Called(ctx, email, code)

	if len(ret) == 0 {
		panic("no return value specified for GetCode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, email, code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCode provides a mock function with given fields: ctx, email, code, expiration
func (_m *AuthMongoRepository) SetCode(ctx context.Context, email string, code int, expiration time.Time) error {
	ret := _m.Called(ctx, email, code, expiration)

	if len(ret) == 0 {
		panic("no return value specified for SetCode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, time.Time) error); ok {
		r0 = rf(ctx, email, code, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAuthMongoRepository creates a new instance of AuthMongoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthMongoRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthMongoRepository {
	mock := &AuthMongoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
