package auth

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"social/internal/domain"
	"time"
)

// Refresh обновляет пару токенов
func (uc *authUseCases) Refresh(ctx context.Context, refreshToken string) (domain.TokenPair, error) {
	if refreshToken == "" {
		err := errors.New("refresh token is required")
		uc.logger.Warn("Empty refresh token", zap.Error(err))
		return domain.TokenPair{}, err
	}

	// проверяем refresh-токен
	userID, err := uc.repo.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		uc.logger.Error("Invalid refresh token", zap.Error(err))
		return domain.TokenPair{}, err
	}

	// генерируем новую пару токенов через репозиторий
	newTokens, err := uc.repo.GenerateTokenPair(userID)
	if err != nil {
		uc.logger.Error("Failed to generate new tokens", zap.Error(err), zap.String("user_id", userID))
		return domain.TokenPair{}, err
	}

	// сохраняем новый refresh-токен
	expiresAt := time.Now().Add(uc.config.JWT.RefreshTokenTTL)
	err = uc.repo.StoreRefreshToken(ctx, userID, newTokens.RefreshToken, expiresAt)
	if err != nil {
		uc.logger.Error("Failed to  store new refresh token", zap.Error(err))
		return domain.TokenPair{}, err
	}

	uc.logger.Info("Tokens refreshed successfully", zap.String("user_id", userID))
	return newTokens, nil
}
