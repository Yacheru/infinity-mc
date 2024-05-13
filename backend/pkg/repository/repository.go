package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yacheru/infinity-mc.ru/backend"
)

type McBans interface {
	GetAllBans(limit int) ([]backend.LbPunishments, error)
}

type Repository struct {
	McBans
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		McBans: NewMcMsql(db),
	}
}
