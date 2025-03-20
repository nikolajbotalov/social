package db

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"social/internal/config"
	"time"
)

// DB представляет подключение к PostgreSQL
type DB struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

// New создает новое подключение к PostgreSQL с повторными попытками
func New(cfgPSQL config.Postgres, logger *zap.Logger) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfgPSQL.Host, cfgPSQL.Port, cfgPSQL.Username, cfgPSQL.Password, cfgPSQL.Database)

	logger.Info("Connecting to PostgreSQL", zap.String("host", cfgPSQL.Host), zap.String("port", cfgPSQL.Port))

	var pool *pgxpool.Pool
	ctx := context.Background()

	err := retry.Do(
		func() error {
			var err error
			pool, err = pgxpool.New(ctx, connStr)
			if err != nil {
				logger.Warn("Failed to create connection pool", zap.Error(err))
				return fmt.Errorf("create pool: %w", err)
			}

			// Проверка соединения через ping
			if err = pool.Ping(ctx); err != nil {
				logger.Warn("Failed to ping db", zap.Error(err))
				pool.Close()
				return fmt.Errorf("ping db: %w", err)
			}

			return nil
		},
		retry.Attempts(10),
		retry.Delay(2*time.Second),
		retry.LastErrorOnly(true),
		retry.OnRetry(func(n uint, err error) {
			logger.Info("Retrying connection", zap.Uint("attempt", n+1), zap.Error(err))
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL after retries: %w", err)
	}

	logger.Info("Successfully connected to PostgreSQL")
	return &DB{
		pool:   pool,
		logger: logger,
	}, nil
}

// Close закрывает пул соединений
func (db *DB) Close() error {
	if db.pool != nil {
		db.pool.Close()
	}
	return nil
}

// Pool возвращает пул соединений для использования в репозиториях
func (db *DB) Pool() *pgxpool.Pool {
	return db.pool
}
