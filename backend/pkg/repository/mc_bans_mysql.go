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

	query := fmt.Sprintf(`
		SELECT ln.uuid, ln.lower_name, lp.reason, lp.start, lp.end,
			CASE 
            	WHEN lp.operator = UNHEX(REPLACE('00000000-0000-0000-0000-000000000000', '-', '')) THEN 'Console'
            	ELSE lp.operator 
       		END AS operator_uuid,
		FROM %s lp
    	INNER JOIN %s lh ON lh.id = lp.id
		INNER JOIN %s lv ON lv.type != 1
    	INNER JOIN %s ln ON lv.uuid = ln.uuid
		LIMIT ?
	`, lbPunishments, lbHistory, lbVictims, lbNames)
	err := mc.db.Select(&bans, query, limit)

	return bans, err
}
