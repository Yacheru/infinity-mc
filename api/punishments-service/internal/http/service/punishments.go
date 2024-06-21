package service

import (
	"punishments-service/internal/repository/mysql"
	"punishments-service/internal/repository/records"
)

type PunishmentService struct {
	repo mysql.Punishment
}

func NewPunishmentService(repo mysql.Punishment) *PunishmentService {
	return &PunishmentService{repo: repo}
}

func (s *PunishmentService) GetPunishments(limit int, pType string) ([]records.LbPunishments, error) {
	return s.repo.GetPunishments(limit, pType)
}
