package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"punishments-service/init/logger"
	"punishments-service/internal/repository/records"
	"punishments-service/pkg/util/constants"
)

type MYSQL struct {
	db *sqlx.DB
}

func NewMySQL(db *sqlx.DB) *MYSQL {
	return &MYSQL{db: db}
}

func (ms *MYSQL) GetPunishments(limit int, pType string) ([]records.LbPunishments, error) {
	var bans []records.LbPunishments
	offset := limit - 10

	logger.Info(pType, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})

	query := fmt.Sprintf(`
		SELECT ln.uuid AS "victim.uuid", COALESCE(ln.name, 'Неизвестно') AS "victim.name", 
			   lp.reason, 
			   lp.start AS "time.start", lp.end AS "time.end", 
			   lp.operator AS "operator.uuid", COALESCE(op.name, 'Консоль') AS "operator.name"
		FROM %s pt 
			INNER JOIN %s lp ON pt.id = lp.id
			INNER JOIN %s lv ON pt.victim = lv.id
			INNER JOIN %s ln ON lv.uuid = ln.uuid
			LEFT JOIN %s op ON lp.operator = op.uuid
		ORDER BY lp.start DESC
		LIMIT ?
		OFFSET ?`,
		constants.LbBans, constants.LbPunishments, constants.LbVictims, constants.LbNames, constants.LbNames,
	)
	err := ms.db.Select(&bans, query, limit, offset)

	for i := range bans {
		bans[i].ID = i + 1
	}

	return bans, err
}
