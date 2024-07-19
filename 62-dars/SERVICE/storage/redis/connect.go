package redisDb

import (
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	return client
}
