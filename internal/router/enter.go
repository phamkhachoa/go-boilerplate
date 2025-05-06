package router

import (
	rauth "go-ecommerce-backend-api/internal/router/auth"
	"go-ecommerce-backend-api/internal/router/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Auth rauth.AuthRouterGroup
}

var RouterGroupApp = new(RouterGroup)
