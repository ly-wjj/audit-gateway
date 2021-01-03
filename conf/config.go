package conf

import (
	"audit-gateway/middleware"
	"audit-gateway/model"
	"github.com/joho/godotenv"
	"os"
)

// You can get the variable from the environment variable, or fill in the required configuration directly in the init function.
func Init() {
	//从本地读取环境变量
	godotenv.Load()
	//数据库初始化
	model.Database(os.Getenv("MYSQL_DSN"))
	//初始化配置
	middleware.SetGoCacheData()

}
