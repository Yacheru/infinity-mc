package services

import (
	"payments-service/internal/repository/psql"
)

type Payment interface {
	CreateHistory(paymentId, nickname, price, donatType string) error
}

type Service struct {
	Payment
}

func NewService(repo *psql.PaymentImpl) *Service {
	return &Service{
		Payment: NewPaymentService(repo.Payment),
	}
}
