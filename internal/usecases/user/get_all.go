package user

import (
	"context"
	"go.uber.org/zap"
	"social/internal/domain"
)

func (uc *useCase) GetAll(ctx context.Context, limit, offset int) ([]domain.User, error) {
	uc.logger.Info("Processing get all users", zap.Int("limit", limit), zap.Int("offset", offset))

	if limit > 100 {
		return nil, domain.ErrLimitExceeded
	}

	if err := ValidatePagination(limit, offset); err != nil {
		uc.logger.Warn("Invalid pagination parameters", zap.Int("limit", limit), zap.Int("offset", offset))
		return nil, err
	}

	users, err := uc.repo.GetAll(ctx, limit, offset)
	if err != nil {
		uc.logger.Error("Failed to fetch users", zap.Error(err))
		return nil, err
	}

	uc.logger.Info("Successfully processed GetAll users", zap.Int("user_count", len(users)))
	return users, nil
}
