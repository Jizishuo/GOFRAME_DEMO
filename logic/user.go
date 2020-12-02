package logic

import (
	"wep_app/dao/mysql"
	"wep_app/pkg/snowflake"
)

func SignUp()  {
	// 1. 判断用户存不存在
	mysql.QueryUserByUserName()
	// 2. 生成uid
	snowflake.GenID()
	// 3. 密码加密

	// 4. 储库
	mysql.InserUser()
}
