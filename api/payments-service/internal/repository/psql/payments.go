package psql

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"payments-service/pkg/util/constants"
)

type PaymentsPSQL struct {
	db *sqlx.DB
}

func NewPaymentsPSQL(db *sqlx.DB) *PaymentsPSQL {
	return &PaymentsPSQL{db: db}
}

func (pPSQL *PaymentsPSQL) CreateHistory(paymentId, nickname, price, donatType string) error {
	tx, err := pPSQL.db.Begin()
	if err != nil {
		return err
	}

	var paymentPK int
	setHistory := fmt.Sprintf("INSERT INTO %s (payment_id, price, donat_type) VALUES ($1, $2, $3) RETURNING id", constants.PaymentsHistory)

	row := tx.QueryRow(setHistory, paymentId, price, donatType)
	err = row.Scan(&paymentPK)
	if err != nil {
		return tx.Rollback()
	}

	setUsersPayments := fmt.Sprintf("INSERT INTO %s (nickname, payment_id) VALUES ($1, $2)", constants.UsersPayments)
	_, err = tx.Exec(setUsersPayments, nickname, paymentPK)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
