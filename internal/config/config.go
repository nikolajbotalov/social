package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"sync"
)

type Config struct {
	BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
	Port   string `env:"PORT" env-default:"8080"`
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
