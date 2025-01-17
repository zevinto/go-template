package route

import (
	"go-template/app/api/admin"

	"github.com/gin-gonic/gin"
)

func genAdminRouter(rg *gin.RouterGroup) {
	rg.GET("/profile", admin.AdminApi.Profile)
}
