package psql

import (
	"fmt"
	"statistic-service"

	"github.com/jmoiron/sqlx"
)

type PaymentsPSQL struct {
	db *sqlx.DB
}

func NewPaymentsPSQL(db *sqlx.DB) *PaymentsPSQL {
	return &PaymentsPSQL{db: db}
}

func (p *PaymentsPSQL) CreateHistory(paymentId, nickname, price, donatType string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	var paymentPK int
	setHistory := fmt.Sprintf("INSERT INTO %s (payment_id, price, donat_type) VALUES ($1, $2, $3) RETURNING id", statistic_service.PaymentsHistory)

	row := tx.QueryRow(setHistory, paymentId, price, donatType)
	err = row.Scan(&paymentPK)
	if err != nil {
		return tx.Rollback()
	}

	setUsersPayments := fmt.Sprintf("INSERT INTO %s (nickname, payment_id) VALUES ($1, $2)", statistic_service.UsersPayments)
	_, err = tx.Exec(setUsersPayments, nickname, paymentPK)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
