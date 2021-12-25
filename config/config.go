package config

import (
	"fmt"

	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

type config struct {
	DB db
}

type db struct {
	Host string `env:"DB_HOST" required:"true"`
	Port int    `env:"DB_PORT" required:"true"`
	Name string `env:"DB_NAME" required:"true"`
	Pass string `env:"DB_PASSWORD"`
	SSL  string `env:"DB_SSL"`
	User string `env:"DB_USER" required:"true"`
}

func InitConfig() (*config, error) {
	conf := &config{}
	if err := configor.Load(conf, "./config/config.yaml"); err != nil {
		return nil, errors.WithStack(err)
	}

	return conf, nil
}

func (d db) CreateDSN() string {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=%s user=%s",
		d.Host,
		d.Port,
		d.Name,
		d.SSL,
		d.User,
	)

	return dsn
}
