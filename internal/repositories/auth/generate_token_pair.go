package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"social/internal/domain"
	"time"
)

func (r *authRepository) GenerateTokenPair(userID string) (domain.TokenPair, error) {
	// Access-токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": jwt.NewNumericDate(time.Now().Add(r.config.JWT.AccessTokenTTL)),
	})
	accessToken, err := token.SignedString(r.config.JWT.Secret)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh-токен
	refreshBytes := make([]byte, 32)
	if _, err = rand.Read(refreshBytes); err != nil {
		return domain.TokenPair{}, fmt.Errorf("failed to generate refresh token: %w", err)
	}
	refreshToken := base64.URLEncoding.EncodeToString(refreshBytes)

	return domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
