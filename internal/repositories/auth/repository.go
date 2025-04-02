package auth

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"social/internal/config"
	"social/internal/domain"
	"time"
)

type Repository interface {
	RegisterUser(ctx context.Context, nickname, password string) error
	Login(cxt context.Context, nickname, password string) (string, domain.TokenPair, error)
	StoreRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) error
	ValidateRefreshToken(ctx context.Context, token string) (string, error)
	GenerateTokenPair(userID string) (domain.TokenPair, error)
}

type authRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
	config *config.Config
}

func NewAuth(db *pgxpool.Pool, logger *zap.Logger, cfg *config.Config) Repository {
	return &authRepository{
		db:     db,
		logger: logger,
		config: cfg,
	}
}
