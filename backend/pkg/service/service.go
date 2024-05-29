package service

import (
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/repository"
)

type McBans interface {
	GetAllBans(limit int) ([]mc.LbPunishments, error)
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
