package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
)

type McBans interface {
	GetAllBans(limit int) ([]mc.LbPunishments, error)
}

type Repository struct {
	McBans
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		McBans: NewMcMsql(db),
	}
}
