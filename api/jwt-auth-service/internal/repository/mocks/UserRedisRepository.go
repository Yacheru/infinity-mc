// Code generated by mockery v2.49.2. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "jwt-auth-service/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// UserRedisRepository is an autogenerated mock type for the UserRedisRepository type
type UserRedisRepository struct {
	mock.Mock
}

// GetUserById provides a mock function with given fields: ctx, userId
func (_m *UserRedisRepository) GetUserById(ctx context.Context, userId string) (*entities.User, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.User, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreNewUser provides a mock function with given fields: ctx, u
func (_m *UserRedisRepository) StoreNewUser(ctx context.Context, u *entities.User) error {
	ret := _m.Called(ctx, u)

	if len(ret) == 0 {
		panic("no return value specified for StoreNewUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRedisRepository creates a new instance of UserRedisRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRedisRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRedisRepository {
	mock := &UserRedisRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
