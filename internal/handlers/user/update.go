package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"social/internal/domain"
	"social/internal/handlers/helpers"
	"social/internal/usecases/user"
)

// Update обновляет информацию о пользователе
func Update(uc user.UseCase, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var updateData domain.User
		if err := c.ShouldBindJSON(&updateData); err != nil {
			logger.Warn("Failed to parse request body", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		err := uc.Update(c.Request.Context(), id, &updateData)
		if err != nil {
			helpers.HandleUserErrors(c, logger, err, id)
			return
		}

		logger.Info("Successfully updated user", zap.String("user_id", id))
		c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
	}
}
