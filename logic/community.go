package logic

import (
	"wep_app/dao/mysql"
	"wep_app/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	// 查数据库Community
	return mysql.GetCommunityList()
}
