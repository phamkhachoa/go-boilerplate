//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service/impl"
)

func InitAuthRouterHandler() (*controller.AuthController, error) {
	wire.Build(
		repo.NewUserRepository,
		impl.NewAuthService,
		controller.NewAuthController)
	return new(controller.AuthController), nil
}
