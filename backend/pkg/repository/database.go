package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Dbname  string
	SSLMode string
}

func NewDatabaseDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mariadb", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Dbname, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
