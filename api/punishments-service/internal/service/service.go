package service

import (
	"punishments-service/internal/entities"
	"punishments-service/internal/repository"
	ps "punishments-service/internal/service/punishments"
)

type PunishmentsService interface {
	GetPunishments(limit int, pType string) ([]entities.LbPunishments, error)
}

type Service struct {
	PunishmentsService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		PunishmentsService: ps.NewPunishmentsService(r.PunishmentsRepository),
	}
}
