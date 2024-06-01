package databases

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/repository"
	"os"
)

func InitDatabases() (*sqlx.DB, *sqlx.DB) {
	mysql, err := repository.NewMySQLDatabaseDB(repository.Config{
		Host:    viper.GetString("db.mdb.host"),
		Port:    viper.GetString("db.mdb.port"),
		User:    viper.GetString("db.mdb.user"),
		Pass:    viper.GetString("db.mdb.pass"),
		Dbname:  viper.GetString("db.mdb.dbname"),
		SSLMode: viper.GetString("db.mdb.sslmode"),
	})

	if err != nil {
		logrus.Errorf("Error connecting to mysql database, %s", err.Error())
		os.Exit(1)
	}

	psql, err := repository.NewPSQLDatabaseDB(repository.Config{
		Host:    viper.GetString("db.psql.host"),
		Port:    viper.GetString("db.psql.port"),
		User:    viper.GetString("db.psql.user"),
		Pass:    viper.GetString("db.psql.pass"),
		Dbname:  viper.GetString("db.psql.dbname"),
		SSLMode: viper.GetString("db.psql.sslmode"),
	})

	if err != nil {
		logrus.Errorf("Error connecting to psql database, %s", err.Error())
		os.Exit(1)
	}

	return mysql, psql
}
