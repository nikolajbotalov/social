package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"social/internal/domain"
)

func (r *userRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.User, error) {
	r.logger.Info("Fetching all users")

	columns := []string{"id", "first_name", "last_name", "nickname", "birthday", "last_visit",
		"interests", "created_at", "updated_at"}
	query, args, err := sq.Select(columns...).From("users").Limit(uint64(limit)).Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		r.logger.Error(
			"Failed to execute query", zap.String("query", query), zap.Any("args", args), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
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
			r.logger.Error("Failed to scan row", zap.Error(err))
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("error while iterating over rows", zap.Error(err))
		return nil, err
	}

	if len(users) == 0 {
		r.logger.Info("No users found")
		return users, nil
	}

	r.logger.Info("Successfully fetched all users", zap.Int("count", len(users)))
	return users, nil
}
