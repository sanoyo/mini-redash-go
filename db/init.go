package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func Init(maxconn int, maxLifetime time.Duration) error {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=%s user=%s",
		"127.0.0.1",
		5433,
		"sample",
		"disable",
		"postgres",
	)

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
