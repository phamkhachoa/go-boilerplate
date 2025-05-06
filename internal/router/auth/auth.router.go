package rauth

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/internal/wire"
)

type AuthRouter struct{}

func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authController, _ := wire.InitAuthRouterHandler()
	userRouterPublic := Router.Group("/auth")
	{
		userRouterPublic.POST("/login", authController.Login)
	}
}
