package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	activePayments  = "active_payments"
	paymentsHistory = "payments_history"
	usersPayments   = "users_payments"
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
		logrus.Errorf("Error starting transaction: %s", err.Error())

		return err
	}

	var paymentPK int
	setHistory := fmt.Sprintf("INSERT INTO %s (payment_id, price, donat_type) VALUES ($1, $2, $3) RETURNING id", paymentsHistory)

	row := tx.QueryRow(setHistory, paymentId, price, donatType)
	err = row.Scan(&paymentPK)
	if err != nil {
		logrus.Errorf("Error adding setHistory transaction: %s", err.Error())

		return tx.Rollback()
	}

	setUsersPayments := fmt.Sprintf("INSERT INTO %s (nickname, payment_id) VALUES ($1, $2)", usersPayments)
	_, err = tx.Exec(setUsersPayments, nickname, paymentPK)
	if err != nil {
		logrus.Errorf("Error add setUsersPayments transaction: %s", err.Error())

		return tx.Rollback()
	}

	return tx.Commit()
}
