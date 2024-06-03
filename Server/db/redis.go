package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := RedisDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
