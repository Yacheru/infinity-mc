package mysql

import (
	"github.com/jmoiron/sqlx"
	"punishments-service/internal/repository/records"
)

type Punishment interface {
	GetPunishments(limit int, pType string) ([]records.LbPunishments, error)
}

type PunishmentsImpl struct {
	Punishment
}

func NewRepository(mysql *sqlx.DB) *PunishmentsImpl {
	return &PunishmentsImpl{
		Punishment: NewMySQL(mysql),
	}
}
