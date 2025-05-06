package service

import (
	"context"
	"go-ecommerce-backend-api/internal/model/dtos"
)

type IAuthService interface {
	Login(ctx context.Context, username string, password string) (codeResult int, out dtos.LoginOutput, error error)
}
