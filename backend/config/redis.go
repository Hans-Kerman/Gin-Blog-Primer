package config

import (
	"log"

	"github.com/Hans-Kerman/GinBlogPrimer/backend/global"
	"github.com/go-redis/redis"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	_, err := RedisClient.Ping().Result() //检测连接

	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	global.RedisDB = RedisClient
}
