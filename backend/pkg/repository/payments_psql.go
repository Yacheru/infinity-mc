package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	activePayments = "active_payments"
)

type PaymentsPSQL struct {
	db *sqlx.DB
}

func NewPaymentsPSQL(db *sqlx.DB) *PaymentsPSQL {
	return &PaymentsPSQL{db: db}
}

func (pPSQL *PaymentsPSQL) AddActivePayment(paymentId string) error {
	query := fmt.Sprintf("INSERT INTO %s (payment_id) VALUES ($1)", activePayments)

	_, err := pPSQL.db.Exec(query, paymentId)

	if err != nil {
		logrus.Errorf("Error adding active payment to payments table: %s", err.Error())
	}

	return nil
}
