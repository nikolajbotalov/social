package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/handlers/helpers"
	"social/internal/usecases/user"
)

// Delete удаление пользователя
func Delete(uc user.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := uc.Delete(c.Request.Context(), id)
		if err != nil {
			helpers.HandleUserErrors(c, logger, err, id)
			return
		}

		logger.Info("Successfully deleted user", zap.String("user_id", id))
		c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
	}
}
