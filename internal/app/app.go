package app

import (
	"go.uber.org/zap"
	"social/internal/config"
	"social/internal/logger"
)

type App struct {
	Logger *zap.Logger
	Config *config.Config
	Server *Server
}

// NewApp инициализирует и возвращает новое приложение
func NewApp() (*App, error) {
	zapLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	cfg := config.LoadConfig(zapLogger)

	// инициализация сервера
	server := NewServer(cfg, zapLogger)

	return &App{
		Logger: zapLogger,
		Config: cfg,
		Server: server,
	}, nil
}

// Close освобождает ресурсы приложения
func (a *App) Close() {
	a.Logger.Info("Closing application")
}
