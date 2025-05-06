package service

import (
	"go-ecommerce-backend-api/core"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/request"
	"go-ecommerce-backend-api/internal/vo"
)

type (
	//.. interfaces
	IUserLogin interface{}
	IUserInfo  interface {
		Register(userRegistratorRequest *vo.UserRegistratorRequest) int
		GetUserById(id int64) (*model.GoCrmUser, int)
		SearchUser(filter *request.UserFilter, paging *core.Paging) ([]model.GoCrmUser, error)
	}
	IUserAdmin interface{}
)
