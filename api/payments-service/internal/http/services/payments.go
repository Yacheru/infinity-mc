package services

import (
	"payments-service/internal/repository/psql"
)

type PaymentService struct {
	repo psql.Payment
}

func NewPaymentService(repo psql.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) CreateHistory(paymentId, nickname, price, donatType string) error {
	return s.repo.CreateHistory(paymentId, nickname, price, donatType)
}
