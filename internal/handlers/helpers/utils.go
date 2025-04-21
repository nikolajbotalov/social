package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/domain"
)

// HandleUserErrors обрабатывает ошибки, связанные с операциями над пользователем
func HandleUserErrors(c *gin.Context, logger *zap.Logger, err error, id string) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		logger.Info("User not found", zap.String("user_id", id))
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	case errors.Is(err, domain.ErrEmptyUserData):
		logger.Info("Empty user data", zap.String("user_id", id))
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty user data"})
	default:
		logger.Error("Failed to process request", zap.String("user_id", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

// GetIntFromContext получение значение int из контекста или возвращает ошибку
func GetIntFromContext(c *gin.Context, key string, logger *zap.Logger) (int, bool) {
	value, exists := c.Get(key)
	if !exists {
		logger.Error("key not set in context", zap.String("key", key))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return 0, false
	}

	valueInt, ok := value.(int)
	if !ok {
		logger.Error("key is not an integer", zap.String("key", key))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return 0, false
	}

	return valueInt, true
}
