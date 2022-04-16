package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/eternalsad/user-auth-service/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDB(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	dbURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s&statement_cache_mode=describe", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SslMode)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()
	pool, err := pgxpool.Connect(ctx, dbURI)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
