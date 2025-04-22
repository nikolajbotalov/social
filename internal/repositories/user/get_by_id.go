package user

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"social/internal/domain"
)

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.logger.Info("Fetching user", zap.String("id", id))

	query, args, err := sq.Select("*").From("users").Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return nil, err
	}

	row := r.db.QueryRow(ctx, query, args...)

	var user domain.User
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Nickname,
		&user.Birthday,
		&user.LastVisit,
		&user.Interests,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("User not found", zap.String("id", id))
			return nil, domain.ErrUserNotFound
		}
		r.logger.Error("Failed to scan row", zap.Error(err))
		return nil, err
	}

	r.logger.Info("Successfully fetched user", zap.String("id", id))
	return &user, nil
}
