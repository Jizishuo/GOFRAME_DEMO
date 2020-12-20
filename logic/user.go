package logic

import (
	"wep_app/dao/mysql"
	"wep_app/models"
	"wep_app/pkg/jwt"
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

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针, 能拿到userid
	if err := mysql.Login(user);err!=nil {
		return "", err
	}
	// 生成jwt
	return jwt.GenToken(user.UserID, user.Username)
}