package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"social/internal/domain"
	"time"
)

func (r *userRepository) Update(ctx context.Context, id string, user *domain.User) error {
	r.logger.Info("Update user", zap.String("id", id))

	user.UpdatedAt = time.Now()

	updateBuilder := sq.Update("users")

	if user.FirstName != nil {
		updateBuilder = updateBuilder.Set("first_name", user.FirstName)
	}
	if user.LastName != nil {
		updateBuilder = updateBuilder.Set("last_name", user.LastName)
	}
	if user.Birthday != nil {
		updateBuilder = updateBuilder.Set("birthday", user.Birthday)
	}
	updateBuilder = updateBuilder.Set("updated_at", user.UpdatedAt)

	query, args, err := updateBuilder.Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build update user query", zap.String("id", id))
		return err
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to update user", zap.String("id", id))
		return err
	}

	// Проверяем, была ли обновлена хотя бы одна строка
	if result.RowsAffected() == 0 {
		r.logger.Error("User not found to update", zap.String("id", id))
		return domain.ErrUserNotFound
	}

	r.logger.Info("Successfully updated user", zap.String("id", id))
	return nil
}
