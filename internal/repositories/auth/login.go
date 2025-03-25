package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"social/internal/domain"
	"time"
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

	accessToken, refreshToken, err := r.generateTokenPair(id)
	if err != nil {
		r.logger.Error("Failed to generate tokens", zap.Error(err), zap.String("id", id))
		return domain.TokenPair{}, err
	}

	r.logger.Info("User logger in successfully")
	return domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *authRepository) generateTokenPair(userID string) (string, string, error) {
	// Access-токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	accessToken, err := token.SignedString(r.jwtSecret)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh-токен
	refreshBytes := make([]byte, 32)
	if _, err = rand.Read(refreshBytes); err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}
	refreshToken := base64.StdEncoding.EncodeToString(refreshBytes)

	return accessToken, refreshToken, nil
}
