package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	lbPunishments = "libertybans_punishments"
	lbBans        = "libertybans_bans"
	lbVictims     = "libertybans_victims"
	lbNames       = "libertybans_names"
	lbHistory     = "libertybans_history"
)

type Config struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Dbname  string
	SSLMode string
}

func NewMySQLDatabaseDB(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Dbname)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPSQLDatabaseDB(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Dbname, cfg.SSLMode)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
