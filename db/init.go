package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func InitDB(maxconn int, maxLifetime time.Duration, dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
