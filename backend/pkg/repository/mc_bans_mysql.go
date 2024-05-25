package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
)

type McMsql struct {
	db *sqlx.DB
}

func NewMcMsql(db *sqlx.DB) *McMsql {
	return &McMsql{db: db}
}

func (mcMsql *McMsql) GetAllBans(limit int) ([]mc.LbPunishments, error) {
	var bans []mc.LbPunishments
	_ = bans

	query := fmt.Sprintf(`
	
	`, lbPunishments, lbHistory, lbVictims, lbNames)
	err := mcMsql.db.Select(&bans, query, limit)

	return bans, err
}
