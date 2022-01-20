package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

var (
	DB  *pgx.Conn
	err error
)

func InitDB(maxconn int, maxLifetime time.Duration, dsn string) error {
	DB, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}

	return nil
}
