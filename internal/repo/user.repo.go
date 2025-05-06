package repo

import (
	"github.com/pkg/errors"
	"go-ecommerce-backend-api/core"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/request"
	"gorm.io/gorm"
	"time"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*model.GoCrmUser, error)
	CreateUser(user *model.GoCrmUser) error
	GetUserById(id int64) (*model.GoCrmUser, error)
	SearchUser(filter *request.UserFilter, paging *core.Paging) ([]model.GoCrmUser, error)
}

type userRepository struct {
	//db *gorm.DB
}

func (ur *userRepository) SearchUser(filter *request.UserFilter, paging *core.Paging) ([]model.GoCrmUser, error) {
	var users []model.GoCrmUser

	db := global.Mdb.Table(TableNameGoCrmUser).Where("1=1")

	if email := filter.Email; email != nil {
		db = db.Where("usr_email = ?", email)
	}

	if phone := filter.Phone; phone != nil {
		db = db.Where("usr_phone = ?", phone)
	}

	// Count total records match conditions
	if err := db.Select("count(*)").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Query data with paging
	if err := db.Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return users, nil
}

func (ur *userRepository) GetUserById(id int64) (*model.GoCrmUser, error) {
	// Thực hiện truy vấn
	var user model.GoCrmUser

	result := global.Mdb.Where("usr_id = ?", id).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Không tìm thấy bản ghi
			return nil, nil
		}
		// Có lỗi khác xảy ra
		return nil, result.Error
	}

	return &user, nil
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) CreateUser(user *model.GoCrmUser) error {
	// Set timestamps
	currentTime := int32(time.Now().Unix())
	user.UsrCreatedAt = currentTime
	user.UsrUpdatedAt = currentTime

	// Set default values if needed
	if user.UsrLoginTimes == 0 {
		user.UsrLoginTimes = 0
	}
	if !user.UsrStatus {
		user.UsrStatus = true // default to enabled
	}

	// Create the user record
	result := global.Mdb.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *userRepository) GetUserByEmail(email string) (*model.GoCrmUser, error) {
	// Thực hiện truy vấn
	var user model.GoCrmUser

	result := global.Mdb.Where("usr_email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Không tìm thấy bản ghi
			return nil, nil
		}
		// Có lỗi khác xảy ra
		return nil, result.Error
	}

	return &user, nil
}

// FindWithFilters retrieves users with applied filters and pagination
func (ur *userRepository) FindWithFilters(filter *request.UserFilter, offset, limit int, sortField string) ([]model.GoCrmUser, error) {
	var users []model.GoCrmUser
	err := global.Mdb.Order(sortField + " " + "desc").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// CountWithFilters counts users with applied filters
func (ur *userRepository) CountWithFilters(query *gorm.DB) (int64, error) {
	var total int64
	err := query.Count(&total).Error
	return total, err
}

// BuildUserFilterQuery applies filters to a GORM query
func (ur *userRepository) BuildUserFilterQuery(filters map[string]interface{}) *gorm.DB {
	query := global.Mdb.Model(&model.GoCrmUser{})

	// Add string equals filters
	for field, value := range filters {
		if strValue, ok := value.(string); ok && strValue != "" {
			switch field {
			case "email":
				query = query.Where("usr_email LIKE ?", "%"+strValue+"%")
			case "phone":
				query = query.Where("usr_phone LIKE ?", "%"+strValue+"%")
			}
		}
	}

	return query
}
