package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/global"
)

func Run() *gin.Engine {
	// LoadConfig
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	return r
	//r.Run(":8002")
}
