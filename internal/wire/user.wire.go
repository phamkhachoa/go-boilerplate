//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service/impl"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		impl.NewUserService,
		controller.NewUserController)
	return new(controller.UserController), nil
}
