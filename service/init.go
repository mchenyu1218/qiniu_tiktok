package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	Rdb *redis.Client
)

// 初始化连接
func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
		PoolSize: 100,      // 连接池大小
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
