package db

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

func InitDB(maxconn int, maxLifetime time.Duration, dsn string) (*sql.DB, error) {
	pool, err := sql.Open("pgx", dsn)
	if err != nil {
		errors.WithStack(err)
	}

	pool.SetMaxIdleConns(maxconn)
	pool.SetMaxOpenConns(maxconn)
	pool.SetConnMaxLifetime(maxLifetime)

	if err := pool.Ping(); err != nil {
		errors.WithStack(err)
	}

	return pool, nil
}
