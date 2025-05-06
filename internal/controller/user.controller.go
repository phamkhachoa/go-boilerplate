package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/core"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/request"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/internal/vo"
	"go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

//type UserController struct {
//	userService *service.UserService
//}
//
//func NewUserController() *UserController {
//	return &UserController{
//		userService: service.NewUserService(),
//	}
//}
//
//func (uc *UserController) GetUserById(c *gin.Context) {
//	//c.JSON(http.StatusOK, response.ResponseData{
//	//	Code:    20001,
//	//	Message: "success",
//	//	Data:    []string{"tipjs", "m10", "anonystick"},
//	//})
//
//	response.SuccessResponse(c, 20001, []string{"tipjs", "m10", "anonystick"})
//}

type UserController struct {
	userService service.IUserInfo
}

func NewUserController(userService service.IUserInfo) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
	}

	global.Logger.Info("Params :", zap.Any("params", params))
	result := uc.userService.Register(&params)
	response.SuccessResponse(c, result, nil)
}

func (uc *UserController) DetailUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	result, _ := uc.userService.GetUserById(int64(id))
	response.SuccessResponse(c, response.ErrCodeSuccess, result)
}

// ShowAccount godoc
// @Summary      Search list User
// @Description  Search list User
// @Tags         users
// // @Accept       json
// // @Produce      json
// // @Param        id   path      int  true  "Account ID"
// // @Success      200  {object}  common.PagedResponse[model.GoCrmUser]
// // @Failure      400  {object}  httputil.HTTPError
// // @Failure      404  {object}  httputil.HTTPError
// // @Failure      500  {object}  httputil.HTTPError
// // @Router       /user/search [post]
func (uc *UserController) SearchUser(c *gin.Context) {
	var filter request.UserFilter
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if err := c.ShouldBindJSON(&filter); err != nil {
		response.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	var paging core.Paging
	paging.Page = page
	paging.Limit = limit

	// Create and return the result
	users, err := uc.userService.SearchUser(&filter, &paging)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, core.SuccessResponse(users, paging, filter))
}
