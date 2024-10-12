package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Redis struct {
		Addr     string
		Db       int
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file :%v", err)
	}

	AppConfig = &Config{}

	// 将读取到的配置信息映射到config结构体中
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Unmarshal config :%v", err)
	}

	// 初始化数据库
	InitDB()
	InitRedis()
}
