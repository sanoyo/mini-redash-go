package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Config *config

type config struct {
	DB *db
}

type db struct {
	Host string `env:"DB_HOST" required:"true"`
	Port int    `env:"DB_PORT" required:"true"`
	Name string `env:"DB_NAME" required:"true"`
	Pass string `env:"DB_PASSWORD"`
	SSL  string `env:"DB_SSL"`
	User string `env:"DB_USER" required:"true"`
}

func InitConfig() {
	_ = configor.Load(&Config, "./config/config.yaml")
}

func (d *db) CreateDSN() string {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=%s user=%s",
		d.Host,
		d.Port,
		d.Name,
		d.SSL,
		d.User,
	)

	return dsn
}
