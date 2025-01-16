package bootstrap

import (
	"fmt"
	"go-template/app/model"
	"go-template/internal/config"
	"go-template/internal/global"
	"go-template/internal/mysql"
)

func init() {
	// 初始化配置文件
	global.Config = config.GetConfig()
	// 初始化数据库
	global.DB = mysql.GetConnection()

	var adminList model.Admin
	global.DB.First(&adminList)
	fmt.Println(adminList)
}
