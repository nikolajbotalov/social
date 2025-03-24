package auth

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (r *authRepository) Login(ctx context.Context, nickname, password string) (string, error) {
	// Формируем SQL-запрос
	query, args, err := sq.Select("id", "nickname", "password_hash").From("auth").
		Where(sq.Eq{"nickname": nickname}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error("Failed to build query", zap.Error(err))
		return "", err
	}

	// Выполняем запрос к БД
	var id, foundNickname, passwordHash string
	err = r.db.QueryRow(ctx, query, args...).Scan(&id, &foundNickname, &passwordHash)
	if err != nil {
		r.logger.Warn("User not found", zap.String("nickname", nickname))
		return "", fmt.Errorf("invalid credentials: %w", err)
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		r.logger.Warn("Password mismatch", zap.String("nickname", nickname))
		return "", fmt.Errorf("invalid credentials: %w", err)
	}

	token, err := r.generateJWT(id)
	if err != nil {
		r.logger.Error("Failed to generate JWT", zap.Error(err), zap.String("id", id))
		return "", err
	}

	r.logger.Info("User logger in successfully")
	return token, nil
}

func (r *authRepository) generateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	tokenString, err := token.SignedString(r.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
