package app

import (
	"go.uber.org/zap"
	"social/internal/adapters/db"
	"social/internal/config"
	"social/internal/logger"
)

type App struct {
	Logger *zap.Logger
	Config *config.Config
	DB     *db.DB
	Server *Server
}

// NewApp инициализирует и возвращает новое приложение
func NewApp() (*App, error) {
	// инициализация логгера
	zapLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	// инициализация конфига
	cfg := config.LoadConfig(zapLogger)
	zapLogger.Info("Config loaded",
		zap.String("db_host", cfg.Postgres.Host),
		zap.String("db_host", cfg.Postgres.Port),
		zap.String("db_host", cfg.Postgres.Username),
		zap.String("db_host", cfg.Postgres.Password))

	// запуск миграция
	if err := db.RunMigrations(cfg.Postgres, zapLogger); err != nil {
		zapLogger.Error("Failed to run migrations", zap.Error(err))
		return nil, err
	}

	// инициализация БД
	dbInstance, err := db.New(cfg.Postgres, zapLogger)
	if err != nil {
		zapLogger.Error("failed to initialize db", zap.Error(err))
		return nil, err
	}

	// инициализация сервера
	server := NewServer(cfg, zapLogger)

	return &App{
		Logger: zapLogger,
		Config: cfg,
		DB:     dbInstance,
		Server: server,
	}, nil
}

// Close освобождает ресурсы приложения
func (a *App) Close() {
	a.Logger.Info("Closing application")
	if err := a.DB.Close(); err != nil {
		a.Logger.Error("Failed to close DB", zap.Error(err))
	}
}
