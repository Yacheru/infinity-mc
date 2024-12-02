package mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"punishments-service/pkg/constants"

	"punishments-service/internal/entities"
)

type RepoPunishments struct {
	db *sqlx.DB
}

func NewPunishmentsRepo(db *sqlx.DB) *RepoPunishments {
	return &RepoPunishments{db: db}
}

func (p *RepoPunishments) GetPunishments(ctx context.Context, limit, punishmentType int) (*[]entities.LbPunishments, error) {
	var bans = new([]entities.LbPunishments)
	offset := limit - 10

	query := fmt.Sprintf(`
		SELECT ln.uuid AS "victim.uuid", COALESCE(ln.name, 'Неизвестно') AS "victim.name", 
			   lp.reason, 
			   lp.start AS "time.start", lp.end AS "time.end", 
			   lp.operator AS "operator.uuid", COALESCE(op.name, 'Консоль') AS "operator.name"
		FROM %s pt 
			INNER JOIN libertybans_punishments lp ON pt.id = lp.id
			INNER JOIN libertybans_victims lv ON pt.victim = lv.id
			LEFT JOIN libertybans_names ln ON lv.uuid = ln.uuid
			LEFT JOIN libertybans_names op ON lp.operator = op.uuid
		WHERE lp.type = %d
		ORDER BY lp.start DESC
		LIMIT ?
		OFFSET ?
	`, constants.LbHistory, punishmentType)
	err := p.db.SelectContext(ctx, bans, query, limit, offset)
	for i := range *bans {
		(*bans)[i].ID = i + 1
	}

	return bans, err
}
