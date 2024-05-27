package service

import (
	"github.com/yacheru/infinity-mc.ru/backend/pkg/repository"
)

type PaymentsService struct {
	repo repository.Payments
}

func NewPaymentsService(repo repository.Payments) *PaymentsService {
	return &PaymentsService{repo: repo}
}

func (s *PaymentsService) AddActivePayment(paymentId string) error {
	return s.repo.AddActivePayment(paymentId)
}
