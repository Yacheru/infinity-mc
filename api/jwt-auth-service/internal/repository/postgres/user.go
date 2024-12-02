package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/entities"
	"jwt-auth-service/pkg/constants"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (a *UserPostgres) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var userEntity = new(entities.User)

	query := `
		SELECT uuid, email, ip, role, password, nickname
		FROM users
		WHERE email = $1;
	`
	if err := a.db.GetContext(ctx, userEntity, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constants.UserNotFoundError
		}
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	return userEntity, nil
}

func (a *UserPostgres) GetUserByRefresh(ctx context.Context, refreshToken string) (*entities.User, error) {
	var user = new(entities.User)

	query := `
		SELECT uuid, email, nickname, role, ip, password
		FROM users
		WHERE refresh_token = $1
	`
	err := a.db.GetContext(ctx, user, query, refreshToken)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	return user, nil
}

func (a *UserPostgres) GetUserID(ctx context.Context, email, password string) (string, error) {
	var uuid string

	query := `
		SELECT uuid
		FROM users 
		WHERE email = $1 AND password = $2
	`
	err := a.db.GetContext(ctx, &uuid, query, email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", constants.UserNotFoundError
		}
		logger.Error(err.Error(), constants.PostgresCategory)
		return "", err
	}

	return uuid, nil
}

func (a *UserPostgres) StoreNewUser(ctx context.Context, u *entities.User) (*entities.User, error) {
	var user = new(entities.User)

	u.Role = constants.Player

	query := `
		INSERT INTO users (uuid, email, ip, password, nickname, role) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING uuid, email, ip, password, nickname, role
	`
	if err := a.db.GetContext(ctx, user, query, u.UserID, u.Email, u.IpAddr, u.Password, u.Nickname, u.Role); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, constants.UserAlreadyExistsError
		}
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	return user, nil
}

func (a *UserPostgres) GetAll(ctx context.Context) (*[]entities.User, error) {
	users := new([]entities.User)

	query := `
		SELECT uuid, nickname, role, email, password, ip
		FROM users
		ORDER BY created_at DESC 
	`
	if err := a.db.SelectContext(ctx, users, query); err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	return users, nil
}

func (a *UserPostgres) DeleteUser(ctx context.Context, id string) error {
	query := `
		DELETE
		FROM users 
		WHERE uuid = $1
	`
	rows, err := a.db.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return err
	}

	count, err := rows.RowsAffected()
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return err
	}

	if count == 0 {
		return constants.UserNotFoundError
	}

	return nil
}

func (a *UserPostgres) UpdateRole(ctx context.Context, id, role string) error {
	query := `
		UPDATE users 
		SET role = $1 
		WHERE uuid = $2
	`
	rows, err := a.db.ExecContext(ctx, query, role, id)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return err
	}

	count, err := rows.RowsAffected()
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return err
	}

	if count == 0 {
		return constants.UserNotFoundError
	}

	return nil
}
