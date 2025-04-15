package user

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"social/internal/domain"
)

func (uc *useCase) Update(cxt context.Context, id string, user *domain.User) error {
	uc.logger.Info("Processing Update user", zap.String("user_id", id))

	if err := ValidateUserID(id); err != nil {
		uc.logger.Warn("Invalid user id", zap.String("user_id", id))
		return err
	}
	if user == nil {
		uc.logger.Warn("Empty user data")
		return domain.ErrEmptyUserData
	}

	err := uc.repo.Update(cxt, id, user)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			uc.logger.Info("User not found for update", zap.String("user_id", id))
			return domain.ErrUserNotFound
		}
		uc.logger.Error("Failed to update user", zap.String("user_id", id), zap.Error(err))
		return err
	}

	uc.logger.Info("Successfully processed Update user", zap.String("user_id", id))
	return nil
}
