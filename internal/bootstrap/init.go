package bootstrap

import (
	"fmt"
	"go-template/internal/config"
	"go-template/internal/global"
	"go-template/internal/logger"
	"go-template/internal/mysql"
)

func init() {
	var err error

	// 初始化配置文件
	global.Config = config.GetConfig()

	// 初始化数据库
	global.DB = mysql.GetConnection()

	// 初始化日志
	if global.Logger, err = logger.New(); err != nil {
		fmt.Println("初始化日志错误：", err)
		return
	}
}
