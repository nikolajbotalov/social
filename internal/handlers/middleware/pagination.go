package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// PaginationMiddleware обрабатывает параметры пагинации,
// проверяет их валидность и устанавливает в контекст в виде int
func PaginationMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitStr := c.DefaultQuery("limit", "10")
		offsetStr := c.DefaultQuery("offset", "0")

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			logger.Warn("Invalid limit parameter", zap.String("limit", limitStr))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
			c.Abort()
			return
		}

		if limit > 100 {
			logger.Warn("Limit exceeds maximum", zap.String("limit", limitStr))
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit exceeded maximum allowed value"})
			c.Abort()
			return
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			logger.Warn("Invalid offset parameter", zap.String("offset", offsetStr))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
			c.Abort()
			return
		}
		c.Set("limit", limit)
		c.Set("offset", offset)
		c.Next()
	}
}
