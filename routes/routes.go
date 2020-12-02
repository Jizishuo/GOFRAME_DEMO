package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wep_app/controllers"
	"wep_app/logger"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.POST("/signup", controllers.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}