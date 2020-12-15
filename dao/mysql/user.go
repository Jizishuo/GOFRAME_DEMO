package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"wep_app/models"
)

const secret = "jizishuo.jizishuo"

var (
	ErrorUserExist = errors.New("用户已经存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

func CheckUserExist(username string) (error) {
	// 查询用户是否存在
	sqlStr := `select count(user_id) from user where username = ?`
	var conut int
	if err := db.Get(&conut, sqlStr, username); err !=nil {
		return err
	}
	if conut > 0 {
		return ErrorUserExist
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

func Login(user *models.User) (err error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `select user_id, username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// 判断密码
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}