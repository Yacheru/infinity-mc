package service

import (
	"context"
	"punishments-service/internal/entities"
	"punishments-service/internal/repository"
)

type PunishmentsService interface {
	GetPunishments(ctx context.Context, limit int, category string) (*[]entities.LbPunishments, error)
}

type Service struct {
	PunishmentsService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		PunishmentsService: NewPunishmentsService(repo.PunishmentsRepository),
	}
}
