package global

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	MdbC   *sql.DB
	Rdb    *redis.Client
)
