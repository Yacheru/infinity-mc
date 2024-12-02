package postgres

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"

	"jwt-auth-service/init/config"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/pkg/constants"
)

func NewPostgresConnection(ctx context.Context, cfg *config.Config, log *logrus.Logger) (*sqlx.DB, error) {
	logger.Debug("connecting to postgres...", constants.PostgresCategory)

	db, err := sqlx.ConnectContext(ctx, "pgx", cfg.PostgresDSN)
	if err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	logger.Info("successfully connect to database. Migrating...", constants.PostgresCategory)

	if err := GooseMigrate(db, log); err != nil {
		logger.Error(err.Error(), constants.PostgresCategory)
		return nil, err
	}

	return db, nil
}

func GooseMigrate(db *sqlx.DB, log *logrus.Logger) error {
	goose.SetLogger(log)

	if err := goose.Up(db.DB, "./migrations"); err != nil {
		return err
	}

	return nil
}
