package logic

import (
	"wep_app/dao/mysql"
	"wep_app/models"
	"wep_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err!=nil {
		return err
	}
	// 2. 生成uid
	usrID := snowflake.GenID()

	// 构造user
	user := models.User{
		UserID: usrID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 密码加密

	// 4. 储库
	err = mysql.InserUser(&user)
	return err
}
