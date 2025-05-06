package user

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/internal/middlewares"
	"go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// this is non-dependency
	//userRepo := repo.NewUserRepository()
	//userService := service.NewUserService(userRepo)
	//userHandlerNonDependency := controller.NewUserController(userService)

	// WIRE go
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
	}

	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(Limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/search", userController.SearchUser)
		userRouterPrivate.GET("/:id", userController.DetailUserById)
	}
}
