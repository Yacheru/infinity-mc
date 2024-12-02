package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"punishments-service/internal/repository/mysql"

	"punishments-service/internal/entities"
)

type PunishmentsRepository interface {
	GetPunishments(ctx context.Context, limit, punishmentType int) (*[]entities.LbPunishments, error)
}

type Repository struct {
	PunishmentsRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PunishmentsRepository: mysql.NewPunishmentsRepo(db),
	}
}
