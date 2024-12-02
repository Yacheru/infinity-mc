package mysql

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"punishments-service/init/config"
	"punishments-service/init/logger"
	"punishments-service/pkg/constants"
)

func InitMySQLDatabase(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	logger.Debug("initializing mysql connection", constants.CategoryMySQL)

	db, err := sqlx.ConnectContext(ctx, "mysql", cfg.MySQLUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Debug("mysql connection initialized", constants.CategoryMySQL)

	return db, nil
}
