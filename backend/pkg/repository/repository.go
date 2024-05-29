package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
)

type McBans interface {
	GetAllBans(limit int) ([]mc.LbPunishments, error)
}

type Payments interface {
	AddActivePayment(paymentId string) error
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
