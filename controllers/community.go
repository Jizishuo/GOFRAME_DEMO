package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// --------------------跟社区相关-----------------------------------

func CommunityHandler(c *gin.Context)  {
	// 查询所有社区分类数据 （comminity_id, community_name） 以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("login.CommunityHandler(),getlist err:", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易服务器错误给前端
		return
	}
	ResponseSuccess(c, data)
}