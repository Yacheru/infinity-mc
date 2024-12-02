package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"payments-service/internal/entities"
	"payments-service/internal/repository/postgres"
)

type PaymentsRepository interface {
	StorePayment(ctx context.Context, payment *entities.PaymentService) error
}

type Repository struct {
	PaymentsRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PaymentsRepository: postgres.NewPayments(db),
	}
}
