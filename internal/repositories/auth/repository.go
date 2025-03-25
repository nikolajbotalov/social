package auth

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"social/internal/domain"
	"time"
)

type Repository interface {
	RegisterUser(ctx context.Context, nickname, password string) error
	Login(cxt context.Context, nickname, password string) (domain.TokenPair, error)
	StoreRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) error
	ValidateRefreshToken(ctx context.Context, token string) (string, error)
}

type authRepository struct {
	db        *pgxpool.Pool
	logger    *zap.Logger
	jwtSecret []byte
}

func NewAuth(db *pgxpool.Pool, logger *zap.Logger, jwtSecret string) Repository {
	return &authRepository{
		db:        db,
		logger:    logger,
		jwtSecret: []byte(jwtSecret),
	}
}
