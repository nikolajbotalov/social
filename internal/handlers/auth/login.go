package auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/usecases/auth"
)

// Login возвращает обработчик для входа пользователя
func Login(uc auth.UseCases, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		var req struct {
			Nickname string `json:"nickname" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.Warn("Invalid request body", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		tokens, err := uc.Login(ctx, req.Nickname, req.Password)
		if err != nil {
			logger.Warn("Login failed", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token":  tokens.AccessToken,
			"refresh_token": tokens.RefreshToken,
		})
	}
}
