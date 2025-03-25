package auth

import (
	"context"
	"go.uber.org/zap"
	"social/internal/config"
	"social/internal/domain"
	authRepository "social/internal/repositories/auth"
)

type UseCases interface {
	Register(ctx context.Context, nickname, password string) error
	Login(ctx context.Context, nickname, password string) (domain.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (domain.TokenPair, error)
}

type authUseCases struct {
	repo   authRepository.Repository
	logger *zap.Logger
	config *config.Config
}

func NewAuthUseCases(repo authRepository.Repository, logger *zap.Logger, cfg *config.Config) UseCases {
	return &authUseCases{
		repo:   repo,
		logger: logger,
		config: cfg,
	}
}
