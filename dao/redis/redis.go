package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"wep_app/settings"
)

// 声明一个全局变量
var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: cfg.Password,
		DB: cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

// 对外封装close 方法
func Close() {
	_ = rdb.Close()
}