package setups

import (
	"github.com/jmoiron/sqlx"

	"punishments-service/init/config"
	"punishments-service/init/mysql"
	"punishments-service/pkg/util/constants"
)

func SetupMysqlDatabase() (*sqlx.DB, error) {
	var dsn string
	switch config.ServerConfig.APIEnvironment {
	case constants.EnvironmentDevelopment:
		dsn = config.ServerConfig.MYSQLDsn
	case constants.EnvironmentProduction:
		dsn = config.ServerConfig.MYSQLURL
	}

	config := mysql.DBConfig{
		DriverName:     config.ServerConfig.MYSQLDriver,
		DataSourceName: dsn,
	}

	conn, err := config.InitMySQLDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
