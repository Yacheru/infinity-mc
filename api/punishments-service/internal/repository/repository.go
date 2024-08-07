package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"punishments-service/internal/entities"
	"punishments-service/internal/repository/mysql/punishments"
)

type PunishmentsRepository interface {
	GetPunishments(ctx *gin.Context, limit int, pType string) ([]entities.LbPunishments, error)
}

type Repository struct {
	PunishmentsRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PunishmentsRepository: punishments.NewPunishmentsRepo(db),
	}
}
