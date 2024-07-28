package punishments

import (
	"punishments-service/internal/entities"
	"punishments-service/internal/repository"
)

type ServicePunishments struct {
	punishments repository.PunishmentsRepository
}

func NewPunishmentsService(punishments repository.PunishmentsRepository) *ServicePunishments {
	return &ServicePunishments{punishments: punishments}
}

func (sp *ServicePunishments) GetPunishments(limit int, pType string) ([]entities.LbPunishments, error) {
	return sp.punishments.GetPunishments(limit, pType)
}
