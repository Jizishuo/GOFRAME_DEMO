package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wep_app/logic"
)

func SignUpHandler(c *gin.Context)  {
	// 1 参数校验

	// 2. 业务处理
	logic.SignUp()
	// 3. 返回响应
	c.JSON(http.StatusOK, "ok")
}
