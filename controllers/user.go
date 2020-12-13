package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"wep_app/logic"
	"wep_app/models"
)

func SignUpHandler(c *gin.Context)  {
	// 1 参数校验
	p := new(models.ParamSignUp)  // 下边传递指针
	// bing 做了判断
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误
		zap.L().Error("siguup with invalid param", zap.Error(err))
		// 判断err 是不是校验器的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
			})
		}

		return
	}
	fmt.Println(p)
	// 手动请求参数校验
	//if len(p.Username) ==0 || len(p.Password) ==0 || len(p.RePassword) == 0 ||p.Password!=p.RePassword {
	//	// 请求参数有误
	//	zap.L().Error("siguup with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}

	// 2. 业务处理
	if err := logic.SignUp(p); err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, "success")
}
