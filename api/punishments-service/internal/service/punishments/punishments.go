package punishments

import (
	"github.com/gin-gonic/gin"
	"punishments-service/internal/entities"
	"punishments-service/internal/repository"
)

type ServicePunishments struct {
	punishments repository.PunishmentsRepository
}

func NewPunishmentsService(punishments repository.PunishmentsRepository) *ServicePunishments {
	return &ServicePunishments{punishments: punishments}
}

func (sp *ServicePunishments) GetPunishments(ctx *gin.Context, limit int, pType string) ([]entities.LbPunishments, error) {
	return sp.punishments.GetPunishments(ctx, limit, pType)
}
