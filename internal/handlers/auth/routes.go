package auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"social/internal/usecases/auth"
)

func SetupAuthRoutes(g *gin.Engine, uc auth.UseCases, logger *zap.Logger) {
	authRoutes := g.Group("api/v1/auth")
	{
		authRoutes.POST("/register", Register(uc, logger))
		authRoutes.POST("/login", Login(uc, logger))
		authRoutes.POST("/refresh", Refresh(uc, logger))
	}
}
