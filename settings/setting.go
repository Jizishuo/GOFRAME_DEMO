package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Conf 全局变量， 保存全局配置信息
var Conf  = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int `mapstructure:"port"`
	Startime     string `mapstructure:"start_time"`
	MachineID    int64 `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level         string `mapstructure:"level"`
	FileName      string `mapstructure:"filename"`
	MaxSize       int `mapstructure:"max_size"`
	MaxAge        int `mapstructure:"max_age"`
	MaxBackups    int `mapstructure:"max_backups"`
}
type MySQLConfig struct {
	Host           string `mapstructure:"host"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	DbName         string `mapstructure:"dbname"`
	Port           int `mapstructure:"port"`
	MaxOpenConns   int `mapstructure:"max_open_conns"`
	MaxIdleConns   int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host           string `mapstructure:"host"`
	Password       string `mapstructure:"password"`
	DB             int `mapstructure:"db"`
	PoolSize       int `mapstructure:"pool_size"`
}


func Init() (err error) {
	// viper.SetConfigFile("./conf/config.yaml")
	// 一般
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf") // 当前路径conf文件夹
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("vippper init err: %v\n", err)
		return
	}
	// 读取配置 反序列化到结构体
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("conf init fail", zap.Error(err))
	}
	// 热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 监听到修改
		zap.L().Info("conf.yaml is change")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("conf init fail", zap.Error(err))
		}
	})
	return nil
}