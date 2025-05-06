package impl

import (
	"context"
	"go-ecommerce-backend-api/internal/model/dtos"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/internal/utils/auth"
	"go-ecommerce-backend-api/pkg/response"
)

type authService struct {
	userRepo repo.IUserRepository
}

func NewAuthService(userRepo repo.IUserRepository) service.IAuthService {
	return &authService{userRepo: userRepo}
}

func (as *authService) Login(ctx context.Context, username string, password string) (codeResult int, out dtos.LoginOutput, error error) {
	// get user by email
	userBase, err := as.userRepo.GetUserByEmail(username)
	if err != nil {
		return response.ErrInvalidUserOrPassword, out, err
	}

	userPass := userBase.UsrPassword
	if userPass != password {
		return response.ErrInvalidUserOrPassword, out, err
	}

	// gen token
	token, err := auth.CreateToken(username)
	if err != nil {
		return 0, dtos.LoginOutput{}, err
	}

	return response.ErrCodeSuccess, dtos.LoginOutput{
		Token:   token,
		Message: "success",
	}, nil
}
