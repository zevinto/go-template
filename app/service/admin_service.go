package service

import (
	"fmt"
	"go-template/app/model"
)

type adminService struct {
	baseService
}

func (a *adminService) Profile() (model model.Admin) {
	if a.db == nil {
		fmt.Println("Error: db is nil in Profile method")
		return model
	}
	a.db.First(&model)
	return model
}
