package repository

import "github.com/jmoiron/sqlx"

type McRepository interface {
}

type Repository struct {
	McRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
