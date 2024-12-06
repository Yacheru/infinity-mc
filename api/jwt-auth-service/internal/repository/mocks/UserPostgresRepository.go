// Code generated by mockery v2.49.2. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "jwt-auth-service/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// UserPostgresRepository is an autogenerated mock type for the UserPostgresRepository type
type UserPostgresRepository struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserPostgresRepository) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *UserPostgresRepository) GetAll(ctx context.Context) (*[]entities.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 *[]entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*[]entities.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *[]entities.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *UserPostgresRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByRefresh provides a mock function with given fields: ctx, refreshToken
func (_m *UserPostgresRepository) GetUserByRefresh(ctx context.Context, refreshToken string) (*entities.User, error) {
	ret := _m.Called(ctx, refreshToken)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByRefresh")
	}

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.User, error)); ok {
		return rf(ctx, refreshToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserID provides a mock function with given fields: ctx, email, password
func (_m *UserPostgresRepository) GetUserID(ctx context.Context, email string, password string) (string, error) {
	ret := _m.Called(ctx, email, password)

	if len(ret) == 0 {
		panic("no return value specified for GetUserID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, email, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreNewUser provides a mock function with given fields: ctx, u
func (_m *UserPostgresRepository) StoreNewUser(ctx context.Context, u *entities.User) (*entities.User, error) {
	ret := _m.Called(ctx, u)

	if len(ret) == 0 {
		panic("no return value specified for StoreNewUser")
	}

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) (*entities.User, error)); ok {
		return rf(ctx, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) *entities.User); ok {
		r0 = rf(ctx, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRole provides a mock function with given fields: ctx, id, role
func (_m *UserPostgresRepository) UpdateRole(ctx context.Context, id string, role string) error {
	ret := _m.Called(ctx, id, role)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserPostgresRepository creates a new instance of UserPostgresRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserPostgresRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserPostgresRepository {
	mock := &UserPostgresRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}