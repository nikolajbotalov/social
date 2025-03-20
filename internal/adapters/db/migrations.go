package db

import (
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"os"
	"social/internal/config"
	"time"
)

// 172.18.0.2:5433

func RunMigrations(cfg config.Postgres, logger *zap.Logger) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	logger.Debug("Attempting to connect to DB", zap.String("dsn", dsn))

	if _, err := os.Stat("migrations"); os.IsNotExist(err) {
		logger.Error("Migrations directory not found")
		return fmt.Errorf("migrations directory not found")
	}
	logger.Debug("Migrations directory found")

	var m *migrate.Migrate
	var err error

	err = retry.Do(
		func() error {
			m, err = migrate.New("file://migrations", dsn)
			if err != nil {
				logger.Warn("DB not ready, retrying...", zap.Error(err))
				return err
			}
			return nil
		},
		retry.Attempts(10),
		retry.Delay(2*time.Second),
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	logger.Info("Migrations applied successfully")
	return nil
}
