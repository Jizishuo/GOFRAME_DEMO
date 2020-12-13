package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"wep_app/models"
)

const secret = "jizishuo.jizishuo"

func CheckUserExist(username string) (error) {
	// 查询用户是否存在
	sqlStr := `select count(user_id) from user where username = ?`
	var conut int
	if err := db.Get(&conut, sqlStr, username); err !=nil {
		return err
	}
	if conut > 0 {
		return errors.New("用户已经存在")
	}
	return nil

}

// InserUser 数据库参入一条记录
func InserUser(user *models.User) (err error) {
	// 密码加密 md5
	user.Password = encryptPassword(user.Password)

	// 执行sql入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(oPasswork string) string {
	h := md5.New()
	// 加盐
	h.Write([]byte(secret))
	h.Sum([]byte(oPasswork))
	return hex.EncodeToString(h.Sum([]byte(oPasswork)))
}