package admin

import (
	"go-template/app/response"
	"go-template/app/service"

	"github.com/gin-gonic/gin"
)

type adminApi struct{}

func (a *adminApi) Profile(c *gin.Context) {
	admin := service.AdminService.Profile()
	response.Success(c, admin)
}
