package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"wep_app/settings"

	// 导入驱动
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	// user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		//viper.GetString("mysql.user"),
		//viper.GetString("mysql.password"),
		//viper.GetString("mysql.host"),
		//viper.GetInt("mysql.port"),
		//viper.GetInt("mysql.dbname"),
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:", zap.Error(err))
		return
	}
	// 最大连接数
	//db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	//db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// 对外封装方法 关闭数据库
func Close()  {
	_ = db.Close()
}