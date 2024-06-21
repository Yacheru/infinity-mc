package setups

import (
	"github.com/jmoiron/sqlx"

	"payments-service/init/config"
	"payments-service/init/psql"
	"payments-service/pkg/util/constants"
)

func SetupPSQLDatabase() (*sqlx.DB, error) {
	var dsn string
	switch config.ServerConfig.APIEnvironment {
	case constants.EnvironmentDevelopment:
		dsn = config.ServerConfig.PSQLDsn
	case constants.EnvironmentProduction:
		dsn = config.ServerConfig.PSQLURL
	}

	config := psql.DBConfig{
		DriverName:     config.ServerConfig.PSQLDriver,
		DataSourceName: dsn,
	}

	conn, err := config.InitPsqlDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
