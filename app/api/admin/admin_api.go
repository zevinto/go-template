package admin

import (
	"fmt"
	"go-template/app/request"
	"go-template/app/response"
	"go-template/app/service"

	"github.com/gin-gonic/gin"
)

type adminApi struct{}

func (a *adminApi) Profile(c *gin.Context) {
	admin := service.AdminService.Profile()
	response.Success(c, admin)
}

func (a *adminApi) Save(c *gin.Context) {
	var admin request.SaveAdmin
	if err := c.ShouldBind(&admin); err != nil {
		fmt.Println(err)
		response.ValidateFailed(c, request.GetErrorMsg(admin, err))
	}
}
