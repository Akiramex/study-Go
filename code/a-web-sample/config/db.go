package config

import (
	"a-web-sample/global"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		AppConfig.Database.Host,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Name,
		AppConfig.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database, got Error: %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to configure database, got Error: %v", err)
	}

	// 赋值给全局变量
	global.Db = db
}
