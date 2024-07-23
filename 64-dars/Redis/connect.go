package redisDB

import (
	"github.com/redis/go-redis/v9"
)

func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
}
