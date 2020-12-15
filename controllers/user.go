package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"wep_app/dao/mysql"
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
			ResponseError(c, CodeInvalidParam)
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
		} else {
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
			//})
			ResponseErrorWithMsg(c, CodeInvalidParam,removeTopStruct(errs.Translate(trans)))
			return
		}
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
		zap.L().Error("logic sigup failed", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "注册失败",
		//})
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	// c.JSON(http.StatusOK, "success")
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context)  {
	// 获取请求参数。校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p);err != nil {
		// 请求参数有误
		zap.L().Error("login with invalid param", zap.Error(err))
		// 判断err 是不是校验器的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
			//})
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}

		// return
	}
	// 落伍逻辑
	if err:=logic.Login(p); err!= nil {
		zap.L().Error("login failed", zap.String("username", p.Username), zap.Error(err))
		//c.JSON(http.StatusOK,gin.H{
		//	"msg": "登录失败",
		//})
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "登录成功",
	//})

	// 返回响应
	ResponseSuccess(c, CodeSuccess)
}