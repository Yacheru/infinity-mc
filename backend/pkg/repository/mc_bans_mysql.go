package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yacheru/infinity-mc.ru/backend"
)

type McMsql struct {
	db *sqlx.DB
}

func NewMcMsql(db *sqlx.DB) *McMsql {
	return &McMsql{db: db}
}

func (mc *McMsql) GetAllBans(limit int) ([]backend.LbPunishments, error) {
	var bans []backend.LbPunishments

	// TODO //
	query := fmt.Sprintf(`
		
	`, lbPunishments, lbHistory, lbVictims, lbNames)
	err := mc.db.Select(&bans, query, limit)

	return bans, err
}
