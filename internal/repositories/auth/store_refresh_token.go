package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

func (r *authRepository) StoreRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) error {
	// Создаем SQL-запрос
	query, args, err := sq.Insert("refresh_tokens").Columns("id", "user_id", "token", "expires_at").
		Values(uuid.New().String(), userID, token, expiresAt).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build refresh token query", zap.Error(err))
		return err
	}

	// Выполняем запрос
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to store refresh token", zap.Error(err), zap.String("user_id", userID))
		return err
	}

	return nil
}
