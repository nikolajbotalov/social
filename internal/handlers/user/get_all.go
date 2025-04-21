package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/domain"
	"social/internal/handlers/helpers"
	"social/internal/usecases/user"
)

// GetAll возвращает список пользователей с пагинацией
func GetAll(uc user.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, ok := helpers.GetIntFromContext(c, "limit", logger)
		if !ok {
			return
		}

		offset, ok := helpers.GetIntFromContext(c, "offset", logger)
		if !ok {
			logger.Error("offset not set in context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		users, err := uc.GetAll(c.Request.Context(), limit, offset)
		if err != nil {
			logger.Error("Failed to fetch users",
				zap.Int("limit", limit),
				zap.Int("offset", offset),
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		logger.Info("Successfully fetched users", zap.Int("user_count", len(users)))
		c.JSON(http.StatusOK, domain.GetAllResponse{
			Users: users,
			Meta: struct {
				Limit  int `json:"limit"`
				Offset int `json:"offset"`
				Total  int `json:"total"`
			}{Limit: limit, Offset: offset, Total: len(users)},
		})
	}
}
