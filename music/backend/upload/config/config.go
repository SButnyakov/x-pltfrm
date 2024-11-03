package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer
	Postgres
	Routes
}

type HTTPServer struct {
	Host string `env:"UPLOAD_HTTP_HOST" env-default:"localhost"`
	Port string `env:"UPLOAD_HTTP_PORT" env-default:"8010"`
}

type Postgres struct {
	Name     string `env:"MUSIC_PG_NAME" env-default:"x-pltfrm/music"`
	User     string `env:"MUSIC_PG_USER" env-default:"postgres"`
	Password string `env:"MUSIC_PG_PASSWORD" env-default:"password"`
	Host     string `env:"MUSIC_PG_HOST" env-default:"localhost"`
	Port     string `env:"MUSIC_PG_PORT" env-defaukt:"5432"`
}

type Routes struct {
	HTTP `yaml:"http" env-required:"true"`
}

type HTTP struct {
	V1 `yaml:"v1" env-required:"true"`
}

type V1 struct {
	Root  string `yaml:"root" env-default:"/v1"`
	Hello string `yaml:"hellp" env-default:"/hello"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	//if err := cleanenv.ReadConfig("./config/routes.yaml", &cfg); err != nil {
	//	return nil, err
	//}

	return &cfg, nil
}
