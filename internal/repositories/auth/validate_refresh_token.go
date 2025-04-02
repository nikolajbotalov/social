package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"time"
)

func (r *authRepository) ValidateRefreshToken(ctx context.Context, token string) (string, error) {
	// Создаем запрос к БД
	query, args, err := sq.Select("user_id").From("refresh_tokens").Where(sq.Eq{"token": token}).
		Where(sq.Gt{"expires_at": time.Now()}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build validate token query", zap.Error(err))
		return "", err
	}

	// Выполняем запрос
	var userID string
	err = r.db.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		r.logger.Warn("Invalid or expired refresh token", zap.String("token", token), zap.Error(err))
		return "", err
	}

	return userID, nil
}
