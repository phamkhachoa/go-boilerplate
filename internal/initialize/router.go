package initialize

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/router"
)

func InitRouter() *gin.Engine {
	//r := gin.Default()

	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	//r.Use() // logging
	//r.Use() // cross
	//r.Use() // limiter
	userRouter := router.RouterGroupApp.User
	authRouter := router.RouterGroupApp.Auth

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/check-status")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		authRouter.InitAuthRouter(MainGroup)
	}
	return r
}
