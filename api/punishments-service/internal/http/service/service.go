package service

import (
	"punishments-service/internal/repository/mysql"
	"punishments-service/internal/repository/records"
)

type Punishment interface {
	GetPunishments(limit int, pType string) ([]records.LbPunishments, error)
}

type Service struct {
	Punishment
}

func NewService(repo *mysql.PunishmentsImpl) *Service {
	return &Service{
		Punishment: NewPunishmentService(repo.Punishment),
	}
}
