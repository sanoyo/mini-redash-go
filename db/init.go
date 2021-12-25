package db

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

func InitDB(maxconn int, maxLifetime time.Duration, dsn string) error {
	pool, err := sql.Open("pgx", dsn)
	if err != nil {
		errors.WithStack(err)
	}
	defer pool.Close()

	pool.SetMaxIdleConns(maxconn)
	pool.SetMaxOpenConns(maxconn)
	pool.SetConnMaxLifetime(maxLifetime)

	if err := pool.Ping(); err != nil {
		errors.WithStack(err)
	}

	return nil
}
