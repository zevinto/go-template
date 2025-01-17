package service

import "gorm.io/gorm"

// service层的基础服务, 其他service服务都继承这个服务
type baseService struct {
	db *gorm.DB
}
