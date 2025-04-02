package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

func (r *authRepository) StoreRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) error {
	// начинаем транзакцию
	tx, err := r.db.Begin(ctx)
	if err != nil {
		r.logger.Error("Failed to begin transaction", zap.Error(err))
		return err
	}

	// откатываем транзакцию в случае ошибки
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				r.logger.Error("Failed to rollback transaction", zap.Error(rollbackErr))
				return
			}
		}
	}()

	// Удаляем старый refresh-токен для данного user_id
	deleteQuery, args, err := sq.Delete("refresh_tokens").Where(sq.Eq{"user_id": userID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build delete query", zap.Error(err))
		return err
	}

	_, err = tx.Exec(ctx, deleteQuery, args...)
	if err != nil {
		r.logger.Error("Failed to delete old token", zap.Error(err))
		return err
	}

	// создаем SQL-запрос
	insertQuery, insertArgs, err := sq.Insert("refresh_tokens").
		Columns("id", "user_id", "token", "expires_at").
		Values(uuid.New().String(), userID, token, expiresAt).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build refresh token query", zap.Error(err))
		return err
	}

	// выполняем запрос
	_, err = r.db.Exec(ctx, insertQuery, insertArgs...)
	if err != nil {
		r.logger.Error("Failed to store refresh token", zap.Error(err), zap.String("user_id", userID))
		return err
	}

	// фиксируем транзакцию
	if err = tx.Commit(ctx); err != nil {
		r.logger.Error("Failed to commit transaction", zap.Error(err))
		return err
	}

	return nil
}
