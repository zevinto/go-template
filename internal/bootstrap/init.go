package bootstrap

import (
	"go-template/internal/config"
	"go-template/internal/global"
)

func init() {
	// 初始化配置文件
	global.Config = config.GetConfig()
}
