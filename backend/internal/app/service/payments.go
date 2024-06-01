package service

import (
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/repository"
)

type PaymentsService struct {
	repo repository.Payments
}

func NewPaymentsService(repo repository.Payments) *PaymentsService {
	return &PaymentsService{repo: repo}
}

func (s *PaymentsService) CreateHistory(nickname, paymentId, price, donatType string) error {
	return s.repo.CreateHistory(nickname, paymentId, price, donatType)
}
