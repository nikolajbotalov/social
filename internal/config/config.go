package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"sync"
)

type Config struct {
	Listen   Listen
	Postgres Postgres
}

type Listen struct {
	BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
	Port   string `env:"PORT" env-default:"8080"`
}

type Postgres struct {
	Username string `env:"PSQL_USERNAME" env-default:"postgres"`
	Password string `env:"PSQL_PASSWORD" env-default:"admin"`
	Host     string `env:"PSQL_HOST" env-default:"host.docker.internal"`
	Port     string `env:"PSQL_PORT" env-default:"5432"`
	Database string `env:"PSQL_DATABASE" env-default:"social"`
}

var instance *Config
var once sync.Once

func LoadConfig(logger *zap.Logger) *Config {
	once.Do(func() {
		logger.Info("loading config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			logger.Error("failed loading config", zap.Error(err))
		}
	})

	logger.Info("config loaded")

	return instance
}
