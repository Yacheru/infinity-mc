package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/pkg/constants"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) SetToken(ctx context.Context, userId string, session *entities.Session) error {
	query := `
		UPDATE users 
		SET refresh_token = $1, expires_in = $2 
		WHERE uuid = $3;
	`
	_, err := p.db.ExecContext(ctx, query, session.RefreshToken, session.ExpiresIn, userId)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return err
	}

	return nil
}

func (p *AuthPostgres) RemoveToken(ctx context.Context, refreshToken string) error {
	query := `
		UPDATE users
		SET refresh_token = null, expires_in = null
		WHERE refresh_token = $1;
	`
	_, err := p.db.ExecContext(ctx, query, refreshToken)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return err
	}

	return nil
}
