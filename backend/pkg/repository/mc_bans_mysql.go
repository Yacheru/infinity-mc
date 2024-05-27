package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
)

type McMySQL struct {
	db *sqlx.DB
}

func NewMcMySQL(db *sqlx.DB) *McMySQL {
	return &McMySQL{db: db}
}

func (mcMSQL *McMySQL) GetAllBans(limit int) ([]mc.LbPunishments, error) {
	var bans []mc.LbPunishments
	_ = bans

	query := fmt.Sprintf(`
	
	`, lbPunishments, lbHistory, lbVictims, lbNames)
	err := mcMSQL.db.Select(&bans, query, limit)

	return bans, err
}
