package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"social/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context, limit, offset int) ([]domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, id string, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewUser(db *pgxpool.Pool, logger *zap.Logger) Repository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
