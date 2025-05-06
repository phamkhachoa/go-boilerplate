package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/model/dtos"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var body dtos.LoginInput
	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
	}

	global.Logger.Info("Body :", zap.Any("body", body))
	code, result, _ := ac.authService.Login(c, body.UserAccount, body.UserPassword)
	response.SuccessResponse(c, code, result)
}
