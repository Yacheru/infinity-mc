package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
)

type McBans interface {
	GetPunishments(limit int, pType string) ([]mc.LbPunishments, error)
}

type Payments interface {
	CreateHistory(nickname, paymentId, price, donatType string) error
}

type Repository struct {
	McBans
	Payments
}

func NewRepository(mySQL, pSQL *sqlx.DB) *Repository {
	return &Repository{
		McBans:   NewMcMySQL(mySQL),
		Payments: NewPaymentsPSQL(pSQL),
	}
}
