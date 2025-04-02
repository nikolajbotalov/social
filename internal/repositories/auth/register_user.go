package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (r *authRepository) RegisterUser(ctx context.Context, nickname, password string) error {
	// Генерируем ID для пользователя
	id := uuid.New().String()

	// Хэшируем пароль
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		r.logger.Error("Failed to hash password", zap.Error(err))
		return err
	}

	// Начинаем транзакцию
	tx, err := r.db.Begin(ctx)
	if err != nil {
		r.logger.Error("Failed to begin transaction", zap.Error(err))
		return err
	}

	// Откатываем транзакцию в случае ошибки
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				r.logger.Error("Failed to rollback transaction", zap.Error(rollbackErr))
			}
			return
		}
	}()

	// Текущая временая метка
	now := time.Now()

	// Формируем SQL-запрос в users
	queryUser, argsUser, err := sq.Insert("users").
		Columns("id", "nickname", "last_visit", "created_at", "updated_at").
		Values(id, nickname, now, now, now).PlaceholderFormat(sq.Dollar).ToSql()

	// Выполняем запрос к БД
	_, err = tx.Exec(ctx, queryUser, argsUser...)
	if err != nil {
		r.logger.Error("Failed to insert into users", zap.Error(err), zap.String("nickname", nickname))
		return err
	}

	// Фомируем SQL-запрос в auth
	queryAuth, argsAuth, err := sq.Insert("auth").Columns("id", "nickname", "password_hash").
		Values(id, nickname, passwordHash).PlaceholderFormat(sq.Dollar).ToSql()

	// Выполняем запрос к БД
	_, err = tx.Exec(ctx, queryAuth, argsAuth...)
	if err != nil {
		r.logger.Error("Failed to insert into auth", zap.Error(err), zap.String("nickname", nickname))
		return err
	}

	// Фиксируем транзакцию
	if err = tx.Commit(ctx); err != nil {
		r.logger.Error("Failed to commit transaction", zap.Error(err))
		return err
	}

	r.logger.Info("User register successfully", zap.String("nickname", nickname), zap.String("id", id))
	return nil
}
