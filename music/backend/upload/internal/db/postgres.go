package db

import (
	"context"
	"fmt"
	"x-pltfrm/music/upload/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGPool(cfg config.Postgres) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return pool, nil
}
