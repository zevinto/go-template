package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 定义一个统一的响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

// Error 失败返回
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func ValidateFailed(c *gin.Context, msg string) {
	Error(c, ValidateError, msg)
}

func PermissionDenied(c *gin.Context) {
	Error(c, HaveNoPermission, ErrorMap[HaveNoPermission])
}

func BusinessFail(c *gin.Context, msg string) {
	Error(c, BusinessError, msg)
}
