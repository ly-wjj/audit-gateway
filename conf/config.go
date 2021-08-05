package conf

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"os"
	"time"

	"audit-gateway/middleware"
	"audit-gateway/model"
)

// You can get the variable from the environment variable, or fill in the required configuration directly in the init function.
func Init() {
	//从本地读取环境变量
	godotenv.Load()
	//数据库初始化
	model.Database(os.Getenv("MYSQL_DSN"))
	// 日志初始化
	LogInit()
	//初始化配置
	conf_name := "route_key"
	middleware.SetGoCacheData(conf_name)
	go middleware.GetIntervalConf(5*time.Second, conf_name)

}

func LogInit() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)
	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.InfoLevel)
}
