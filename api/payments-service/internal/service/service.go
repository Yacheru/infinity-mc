package service

import (
	"context"
	"payments-service/internal/entities"
	"payments-service/internal/repository"
	"payments-service/internal/service/payments"
)

type PaymentsService interface {
	CreatePayment(ctx context.Context, payment *entities.PaymentService) (*entities.Payment, error)
	AcceptPayment(ctx context.Context, paid *entities.Paid) error
}

type Service struct {
	PaymentsService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		PaymentsService: payments.NewPayments(repo.PaymentsRepository),
	}
}
