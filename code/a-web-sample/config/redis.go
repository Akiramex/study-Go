package config

import (
	"a-web-sample/global"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr,
		DB:       AppConfig.Redis.Db,
		Password: AppConfig.Redis.Password,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connnect to Redis, got error: %v", err)
	}

	global.RedisDb = RedisClient
}
