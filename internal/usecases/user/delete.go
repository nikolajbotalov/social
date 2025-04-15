package user

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"social/internal/domain"
)

func (uc *useCase) Delete(ctx context.Context, id string) error {
	uc.logger.Info("Processing Delete user", zap.String("user_id", id))

	if err := ValidateUserID(id); err != nil {
		uc.logger.Warn("Invalid user id", zap.String("user_id", id))
		return err
	}

	err := uc.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			uc.logger.Warn("User not found for deletion", zap.String("user_id", id))
			return domain.ErrUserNotFound
		}
		uc.logger.Error("Failed to delete user", zap.String("user_id", id), zap.Error(err))
		return err
	}

	uc.logger.Info("Successfully processed Delete user", zap.String("user_id", id))
	return nil
}
