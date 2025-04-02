package auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/usecases/auth"
)

// Register возвращает обработчик для регистрации пользователя
func Register(uc auth.UseCases, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var req struct {
			Nickname string `json:"nickname" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.Warn("Failed request body", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		err := uc.Register(ctx, req.Nickname, req.Password)
		if err != nil {
			logger.Warn("Registration failed", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
	}
}
