package service

import (
	"payments-service/internal/repository/psql"
)

type Payment interface {
}

type Service struct {
	Payment
}

func NewService(repo *psql.PaymentImpl) *Service {
	return &Service{
		Payment: NewPaymentService(repo.Payment),
	}
}
