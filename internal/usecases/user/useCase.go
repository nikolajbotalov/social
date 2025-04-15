package user

import (
	"context"
	"go.uber.org/zap"
	"social/internal/domain"
	userRepository "social/internal/repositories/user"
)

type UseCase interface {
	GetAll(ctx context.Context, limit, offset int) ([]domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, id string, user *domain.User) error
	Delete(cxt context.Context, id string) error
}

type useCase struct {
	repo   userRepository.Repository
	logger *zap.Logger
}

func NewUserUseCases(repo userRepository.Repository, logger *zap.Logger) UseCase {
	return &useCase{
		repo:   repo,
		logger: logger,
	}
}
