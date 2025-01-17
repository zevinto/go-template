package service

import (
	_ "go-template/internal/bootstrap"
	"go-template/internal/global"
)

var (
	BaseService = &baseService{db: global.DB}

	AdminService = &adminService{baseService: *BaseService}
)
