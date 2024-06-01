package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	LbPunishments = "libertybans_punishments"
	LbBans        = "libertybans_bans"
	LbMutes       = "libertybans_mutes"
	LbWarns       = "libertybans_warns"
	LbVictims     = "libertybans_victims"
	LbNames       = "libertybans_names"
	LbHistory     = "libertybans_history"
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
