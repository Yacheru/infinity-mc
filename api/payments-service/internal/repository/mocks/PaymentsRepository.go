// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "payments-service/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// PaymentsRepository is an autogenerated mock type for the PaymentsRepository type
type PaymentsRepository struct {
	mock.Mock
}

// StorePayment provides a mock function with given fields: ctx, payment
func (_m *PaymentsRepository) StorePayment(ctx context.Context, payment *entities.PaymentService) error {
	ret := _m.Called(ctx, payment)

	if len(ret) == 0 {
		panic("no return value specified for StorePayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.PaymentService) error); ok {
		r0 = rf(ctx, payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPaymentsRepository creates a new instance of PaymentsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentsRepository {
	mock := &PaymentsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
