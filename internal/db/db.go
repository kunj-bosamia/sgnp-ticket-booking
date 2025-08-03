package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kunjbosamia/sgnp-ticket-booking/internal/config"
)

var Pool *pgxpool.Pool

func InitDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.AppConfig.DBUser,
		config.AppConfig.DBPass,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse pool config from DSN
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse DB config: %v", err)
	}

	// Set pool sizing from config
	cfg.MaxConns = int32(config.AppConfig.DBPoolMaxConns)
	cfg.MinConns = int32(config.AppConfig.DBPoolMinConns)

	// Create connection pool with custom config
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Unable to create DB pool: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping DB: %v", err)
	}

	log.Println("Connected to DB")
	Pool = pool
}
