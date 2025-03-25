package auth

import (
	"context"
	"go.uber.org/zap"
)

// Register регистрирует нового пользователя
func (uc *authUseCases) Register(ctx context.Context, nickname, password string) error {
	// валидация входных данных
	if err := ValidateCredentials(nickname, password); err != nil {
		uc.logger.Warn("Invalid registration data", zap.Error(err))
		return err
	}

	// вызов репозитория для регистрации
	err := uc.repo.RegisterUser(ctx, nickname, password)
	if err != nil {
		uc.logger.Error("Failed to register user", zap.Error(err), zap.String("nickname", nickname))
		return err
	}

	uc.logger.Info("User registered successfully", zap.String("nickname", nickname))
	return nil
}
