package database

import (
	"context"
	"fmt"
	"log"
	"teamtask-api/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&channel_binding=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
		cfg.ChannelBinding,
	)

	// create Configuration from DSN
	pgxCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	pgxCfg.MaxConns = 10
	pgxCfg.MinConns = 2
	pgxCfg.MaxConnLifetime = 30 * time.Minute
	pgxCfg.MaxConnIdleTime = 15 * time.Minute

	// Connect to database pool
	ctx, cancle := context.WithTimeout(ctx, 10*time.Second)
	defer cancle()

	pool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		return nil, fmt.Errorf("unnable to connect to database: %w", err)
	}
	// check connection database
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unnable to ping database: %w", err)
	}

	log.Printf("Connected to postgres database at %s:%d/%s", cfg.Host, cfg.Port, cfg.DBName)
	return pool, nil

}
