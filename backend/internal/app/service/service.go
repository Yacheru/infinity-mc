package service

import (
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/repository"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
)

type McBans interface {
	GetPunishments(limit int, pType string) ([]mc.LbPunishments, error)
}

type Payments interface {
	CreateHistory(nickname, paymentId, price, donatType string) error
}

type Service struct {
	McBans
	Payments
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		McBans:   NewMcBansService(repo.McBans),
		Payments: NewPaymentsService(repo.Payments),
	}
}
