package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/handlers/helpers"
	"social/internal/usecases/user"
)

// GetByID возвращает пользователя по его ID
func GetByID(uc user.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		userData, err := uc.GetByID(c.Request.Context(), id)
		if err != nil {
			helpers.HandleUserErrors(c, logger, err, id)
			return
		}

		logger.Info("Successfully fetched user", zap.String("user_id", id))
		c.JSON(http.StatusOK, userData)
	}
}
