package psql

import (
	"github.com/jmoiron/sqlx"
)

type Payment interface {
	CreateHistory(paymentId, nickname, price, donatType string) error
}

type PaymentImpl struct {
	Payment
}

func NewRepository(mysql *sqlx.DB) *PaymentImpl {
	return &PaymentImpl{
		Payment: NewPaymentsPSQL(mysql),
	}
}
