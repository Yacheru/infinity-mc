package service

import (
	"context"

	"punishments-service/init/logger"
	"punishments-service/internal/entities"
	"punishments-service/internal/repository"
	"punishments-service/pkg/constants"
)

type Punishments struct {
	punishments repository.PunishmentsRepository
}

func NewPunishmentsService(punishments repository.PunishmentsRepository) *Punishments {
	return &Punishments{punishments: punishments}
}

func (p *Punishments) GetPunishments(ctx context.Context, limit int, category string) (*[]entities.LbPunishments, error) {
	var punishmentType int

	switch category {
	case constants.Bans:
		punishmentType = 0
	case constants.Mutes:
		punishmentType = 1
	case constants.Warns:
		punishmentType = 2
	}

	bans, err := p.punishments.GetPunishments(ctx, limit, punishmentType)
	if err != nil {
		logger.Error(err.Error(), constants.CategoryService)
		return nil, err
	}
	return bans, nil
}
