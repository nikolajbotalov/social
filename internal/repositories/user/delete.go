package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"social/internal/domain"
)

func (r *userRepository) Delete(ctx context.Context, id string) error {
	r.logger.Info("Remove user", zap.String("id", id))

	query, args, err := sq.Delete("users").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build delete user query", zap.String("id", id))
		return err
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		r.logger.Error("Failed to delete user", zap.String("id", id))
		return err
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("User not found to delete", zap.String("id", id))
		return domain.ErrUserNotFound
	}

	r.logger.Info("Successfully deleted user", zap.String("id", id))
	return nil
}
