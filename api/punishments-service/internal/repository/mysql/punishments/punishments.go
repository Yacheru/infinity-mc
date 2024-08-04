package punishments

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"punishments-service/internal/entities"
	"punishments-service/pkg/util/constants"
)

type RepoPunishments struct {
	db *sqlx.DB
}

func NewPunishmentsRepo(db *sqlx.DB) *RepoPunishments {
	return &RepoPunishments{db: db}
}

func (rp *RepoPunishments) GetPunishments(ctx *gin.Context, limit int, pType string) ([]entities.LbPunishments, error) {
	var bans []entities.LbPunishments
	offset := limit - 10

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
	err := rp.db.SelectContext(ctx.Request.Context(), &bans, query, limit, offset)

	for i := range bans {
		bans[i].ID = i + 1
	}

	return bans, err
}
