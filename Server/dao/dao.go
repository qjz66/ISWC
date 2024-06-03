package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao struct {
	*gorm.DB
}

type Redis struct {
	*redis.Client
}

// NewDao 创建一个Dao实例
func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db.Session(&gorm.Session{}),
	}
}

// NewRedis 创建一个Redis实例
func NewRedis(redis *redis.Client) *Redis {
	return &Redis{redis}
}
