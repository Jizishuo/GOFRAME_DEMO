package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"wep_app/middlewares"
)

var (
	ErrorUserNotLogin = errors.New("用户未登录")
)

// 获取当前登录用户的id
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.CtxtUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
