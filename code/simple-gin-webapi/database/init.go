package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	dbName := viper.GetString("db.name")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host,
		username,
		password,
		dbName,
		port,
	)
	config := gorm.Config{}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		fmt.Println("数据库连接失败")
	} else {
		fmt.Println("数据库链接成功")
	}
}
