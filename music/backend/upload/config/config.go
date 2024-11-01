package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer
	Postgres
}

type HTTPServer struct {
	Host string `env:"HTTP_HOST" env-default:"localhost"`
	Port string `env:"HTTP_PORT" env-default:"8010"`
}

type Postgres struct {
	Name     string `env:"DB_NAME" env-default:"x-pltfrm/music"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"password"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-defaukt:"5432"`
}

func Load() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
