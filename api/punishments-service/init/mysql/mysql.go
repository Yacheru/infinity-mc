package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	DriverName     string
	DataSourceName string
}

func (config *DBConfig) InitMySQLDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %s", err.Error())
	}

	return db, nil
}
