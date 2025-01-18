package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP     HTTP
		Postgres Postgres
		CoinAPI  CoinAPI
		Observer Observer
	}

	HTTP struct {
		Port           string `env:"HTTP_PORT" envDefault:"8080"`
		StartTimeout   int    `env:"HTTP_START_TIMEOUT" envDefault:"5"`
		MaxHeaderBytes int    `env:"HTTP_MAX_HEADER_BYTES" envDefault:"134217728"`
		ReadTimeout    int    `env:"HTTP_READ_TIMEOUT" envDefault:"10"`
		WriteTimeout   int    `env:"HTTP_WRITE_TIMEOUT" envDefault:"10"`
	}

	Postgres struct {
		Host     string `env:"POSTGRES_HOST" envDefault:"127.0.0.1"`
		Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
		Username string `env:"POSTGRES_USER" envDefault:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
		Database string `env:"POSTGRES_DB" envDefault:"postgres"`
	}

	CoinAPI struct {
		Token string `env:"API_KEY"`
	}

	Observer struct {
		ObserveTime int `env:"OBSERVE_TIME"`
	}
)

func New() (*Config, error) {
	godotenv.Load(".env")

	cfg := &Config{}

	err := env.Parse(cfg)
	return cfg, err
}

func (c *Config) DBInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Postgres.Host, c.Postgres.Port, c.Postgres.Username, c.Postgres.Password, c.Postgres.Database)
}
