package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

func InitDB(maxconn int, maxLifetime time.Duration, dsn string) (*pgx.Conn, error) {
	// pool, err := sql.Open("pgx", dsn)
	// if err != nil {
	// 	errors.WithStack(err)
	// }

	// pool.SetMaxIdleConns(maxconn)
	// pool.SetMaxOpenConns(maxconn)
	// pool.SetConnMaxLifetime(maxLifetime)

	// if err := pool.Ping(); err != nil {
	// 	errors.WithStack(err)
	// }
	// return pool, nil

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
