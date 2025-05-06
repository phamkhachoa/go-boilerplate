package impl

import (
	"go-ecommerce-backend-api/core"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/request"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/internal/vo"
	"go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
	"time"
)

//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{
//		userRepo: repo.NewUserRepo(),
//	}
//}
//
//func (us *UserService) GetInfoUser() string {
//	return us.userRepo.GetInfoUser()
//}

//type IUserService interface {
//	Register(userRegistratorRequest *vo.UserRegistratorRequest) int
//	GetUserById(id int64) (*model.GoCrmUser, int)
//	SearchUser(filter *request.UserFilter, page int, limit int) common.PagedResponse[model.GoCrmUser]
//}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func (us *userService) SearchUser(filter *request.UserFilter, paging *core.Paging) ([]model.GoCrmUser, error) {
	// Build filter map to pass to repository

	//var user []model.GoCrmUser
	//global.Mdb.Raw("SELECT * FROM go_crm_user WHERE usr_email = @email OR usr_phone = @phone",
	//	sql.Named("email", filter.Email), sql.Named("phone", filter.Phone)).Find(&user)
	//
	//var totalCount int
	//
	//// For GORM v2
	//global.Mdb.Raw(
	//	"SELECT count(*) FROM go_crm_user WHERE usr_email = @email OR usr_phone = @phone",
	//	sql.Named("email", filter.Email),
	//	sql.Named("phone", filter.Phone),
	//).Scan(&totalCount)
	//
	//return common.NewPagedResponse(
	//	true,
	//	"Products retrieved successfully",
	//	user,
	//	page,
	//	limit,
	//	totalCount,
	//)

	// todo validate

	// call repo
	users, err := us.userRepo.SearchUser(filter, paging)
	return users, err
}

func (us *userService) GetUserById(id int64) (*model.GoCrmUser, int) {
	user, err := us.userRepo.GetUserById(id)
	if err != nil {
		global.Logger.Error("Get Detail User Error", zap.Any("err", err))
	}

	if user == nil {
		global.Logger.Error("User Not Found", zap.Any("err", err))
		return nil, response.ErrUserNotFound
	}

	return user, response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) service.IUserInfo {
	return &userService{userRepo: userRepo, userAuthRepo: userAuthRepo}
}

func (us *userService) Register(userRegistratorRequest *vo.UserRegistratorRequest) int {
	// 1. verify email existed or not
	email := userRegistratorRequest.Email
	//password := userRegistratorRequest.Password

	user, err := us.userRepo.GetUserByEmail(email)

	if err != nil || user != nil {
		global.Logger.Error("User Repository GetUserByEmail Error", zap.Error(err))
		return response.ErrCodeUserHasExists
	}

	// 2. save user
	var goCrmUser model.GoCrmUser

	goCrmUser = model.GoCrmUser{
		UsrEmail:     userRegistratorRequest.Email,
		UsrPhone:     userRegistratorRequest.Phone,
		UsrPassword:  userRegistratorRequest.Password,
		UsrUsername:  email,
		UsrCreatedAt: int32(time.Now().UnixMilli()),
		UsrUpdatedAt: int32(time.Now().UnixMilli()),
		UsrStatus:    false,
	}

	err = us.userRepo.CreateUser(&goCrmUser)
	if err != nil {
		return response.ErrCodeUserHasExists
	}

	//// 0. hashEmail
	//hashEmail := crypto.GetHash(email)
	//
	//if us.userRepo.GetUserByEmail(email) {
	//	return response.ErrCodeUserHasExists
	//}
	//
	//otp := random.GenerateSixDigitOtp()
	//if purpose == "TEST_USER" {
	//	otp = 123456
	//}
	//
	//global.Logger.Info("Otp is::: %d", zap.Int("otp", otp))
	//
	//// save OTP in Redis with expiration time
	//err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	//
	//if err != nil {
	//	return response.ErrInvalidOTP
	//}
	//
	//err = sendto.SendTemplateEmailOtp([]string{email}, "phamkhachoabk@gmail.com", "otp-auth.html", map[string]interface{}{
	//	"otp": otp,
	//})
	//if err != nil {
	//	return response.ErrSendEmailOtp
	//}
	//
	return response.ErrCodeSuccess
}
