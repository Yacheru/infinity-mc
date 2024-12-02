package postgres

import (
	"context"
	"payments-service/pkg/constants"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v5/stdlib"
	"payments-service/init/config"
	"payments-service/init/logger"
)

func InitPostgresqlConnection(ctx context.Context, cfg *config.Config, log *logrus.Logger) (*sqlx.DB, error) {
	logger.Debug("connecting to postgres...", constants.LoggerCategoryPostgres)

	db, err := sqlx.ConnectContext(ctx, "pgx", cfg.PostgresDSN)
	if err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryPostgres)
		return nil, err
	}
	logger.Debug("successfully connect to database. Migrating...", constants.LoggerCategoryPostgres)

	if err := GooseMigrate(db, log); err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryPostgres)

		return nil, err
	}

	logger.Info("postgres is working", constants.LoggerCategoryPostgres)

	return db, nil
}

func GooseMigrate(db *sqlx.DB, log *logrus.Logger) error {
	goose.SetLogger(log)

	if err := goose.Up(db.DB, "./migrations"); err != nil {
		return err
	}

	return nil
}
