package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"social/internal/handlers/middleware"
	"social/internal/usecases/user"
)

func SetupUserRoutes(g *gin.Engine, uc user.UseCase, logger *zap.Logger) {
	userRoutes := g.Group("api/v1/user")
	{
		userRoutes.GET("/all", middleware.PaginationMiddleware(logger), GetAll(uc, logger))
		userRoutes.GET("/:id", middleware.ValidateUserID(logger), GetByID(uc, logger))
		userRoutes.PATCH("/:id", middleware.ValidateUserID(logger), Update(uc, logger))
		userRoutes.DELETE("/:id", middleware.ValidateUserID(logger), Delete(uc, logger))
	}
}
