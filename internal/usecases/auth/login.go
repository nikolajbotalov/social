package auth

import (
	"context"
	"go.uber.org/zap"
	"social/internal/domain"
	"time"
)

// Login выполняет вход пользователя
func (uc *authUseCases) Login(ctx context.Context, nickname, password string) (domain.TokenPair, error) {
	// валидация входных данных
	if err := ValidateCredentials(nickname, password); err != nil {
		uc.logger.Warn("Invalid login data", zap.Error(err))
		return domain.TokenPair{}, err
	}

	// выполняет вход
	userID, tokens, err := uc.repo.Login(ctx, nickname, password)
	if err != nil {
		uc.logger.Error("Failed login", zap.Error(err), zap.String("nickname", nickname))
		return domain.TokenPair{}, err
	}

	// Сохраняем refresh-токен
	expiresAt := time.Now().Add(uc.config.JWT.AccessTokenTTL)
	err = uc.repo.StoreRefreshToken(ctx, userID, tokens.RefreshToken, expiresAt)
	if err != nil {
		uc.logger.Error("Failed to store refresh token", zap.Error(err))
		return domain.TokenPair{}, err
	}

	uc.logger.Info("User logged in successfully", zap.String("nickname", nickname))
	return tokens, nil
}
