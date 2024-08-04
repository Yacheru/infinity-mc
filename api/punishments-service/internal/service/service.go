package service

import (
	"github.com/gin-gonic/gin"
	"punishments-service/internal/entities"
	"punishments-service/internal/repository"
	"punishments-service/internal/service/punishments"
)

type PunishmentsService interface {
	GetPunishments(ctx *gin.Context, limit int, pType string) ([]entities.LbPunishments, error)
}

type Service struct {
	PunishmentsService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		PunishmentsService: punishments.NewPunishmentsService(repo.PunishmentsRepository),
	}
}
