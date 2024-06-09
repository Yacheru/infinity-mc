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

func (mcMSQL *McMySQL) GetPunishments(limit int, pType string) ([]mc.LbPunishments, error) {
	var bans []mc.LbPunishments
	offset := limit - 10

	query := fmt.Sprintf(`
		SELECT ln.uuid AS "victim.uuid", COALESCE(ln.name, 'Неизвестно') AS "victim.name", 
			   lp.reason, 
			   lp.start AS "time.start", 
			   lp.end AS "time.end", 
			   lp.operator AS "operator.uuid", 
			   COALESCE(op.name, 'Console') AS "operator.name"
		FROM %s pt 
			INNER JOIN %s lp ON pt.id = lp.id
			INNER JOIN %s lv ON pt.victim = lv.id
			INNER JOIN %s ln ON lv.uuid = ln.uuid
			LEFT JOIN %s op ON lp.operator = op.uuid
		ORDER BY lp.start DESC
		LIMIT ?
		OFFSET ?`,
		pType, LbPunishments, LbVictims, LbNames, LbNames,
	)
	err := mcMSQL.db.Select(&bans, query, limit, offset)

	for i := range bans {
		bans[i].ID = i + 1
	}

	return bans, err
}
