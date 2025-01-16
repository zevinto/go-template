package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppRouter struct{}

func (*AppRouter) AddRoute(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ping"})
	})

	e.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func New() *AppRouter {
	return &AppRouter{}
}
