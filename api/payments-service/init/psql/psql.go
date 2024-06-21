package psql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	DriverName     string
	DataSourceName string
}

func (config *DBConfig) InitPsqlDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %s", err.Error())
	}

	return db, nil
}
