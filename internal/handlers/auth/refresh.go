package auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/usecases/auth"
)

// Refresh возвращает обработчик для обновления токенов
func Refresh(uc auth.UseCases, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		var req struct {
			RefreshToken string `json:"refresh_token" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.Warn("Invalid request body", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		newTokens, err := uc.Refresh(ctx, req.RefreshToken)
		if err != nil {
			logger.Warn("Token refresh failed", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_tokens": newTokens.AccessToken,
			"refresh_token": newTokens.RefreshToken,
		})
	}
}
