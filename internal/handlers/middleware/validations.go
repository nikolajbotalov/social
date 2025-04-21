package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func ValidateUserID(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			logger.Warn("Invalid user ID", zap.String("user_id", id))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			c.Abort()
			return
		}
		c.Next()
	}
}
