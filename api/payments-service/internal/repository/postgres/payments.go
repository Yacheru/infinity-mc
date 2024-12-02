package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"payments-service/internal/entities"
)

type Payments struct {
	db *sqlx.DB
}

func NewPayments(db *sqlx.DB) *Payments {
	return &Payments{db: db}
}

func (p *Payments) StorePayment(ctx context.Context, payment *entities.PaymentService) error {
	query := `
	 	INSERT INTO payments (email, nickname, service, price, duration) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := p.db.ExecContext(ctx, query, payment.Email, payment.Nickname, payment.Service, payment.Price, payment.Duration)
	if err != nil {
		return err
	}

	return nil
}
