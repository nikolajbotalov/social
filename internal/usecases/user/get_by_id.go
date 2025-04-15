package user

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"social/internal/domain"
)

func (uc *useCase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	uc.logger.Info("Processing GetByID request", zap.String("user_id", id))

	if err := ValidateUserID(id); err != nil {
		uc.logger.Warn("Invalid user ID", zap.String("user_id", id))
		return nil, err
	}

	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			uc.logger.Warn("User not found", zap.String("user_id", id))
			return nil, domain.ErrUserNotFound
		}
		uc.logger.Error("Failed to get user", zap.Error(err), zap.String("user_id", id))
		return nil, err
	}

	uc.logger.Info("Successfully processed GetByID", zap.String("user_id", id))
	return user, nil
}
