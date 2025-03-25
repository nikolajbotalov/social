package auth

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"social/internal/domain"
)

func (r *authRepository) Login(ctx context.Context, nickname, password string) (domain.TokenPair, error) {
	// Формируем SQL-запрос
	query, args, err := sq.Select("id", "nickname", "password_hash").From("auth").
		Where(sq.Eq{"nickname": nickname}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return domain.TokenPair{}, err
	}

	// Выполняем запрос к БД
	var id, foundNickname, passwordHash string
	err = r.db.QueryRow(ctx, query, args...).Scan(&id, &foundNickname, &passwordHash)
	if err != nil {
		r.logger.Warn("User not found", zap.String("nickname", nickname))
		return domain.TokenPair{}, fmt.Errorf("invalid credentials: %w", err)
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		r.logger.Warn("Password mismatch", zap.String("nickname", nickname))
		return domain.TokenPair{}, fmt.Errorf("invalid credentials: %w", err)
	}

	tokens, err := r.GenerateTokenPair(id)
	if err != nil {
		r.logger.Error("Failed to generate tokens", zap.Error(err), zap.String("id", id))
		return domain.TokenPair{}, err
	}

	r.logger.Info("User logger in successfully")
	return tokens, nil
}
