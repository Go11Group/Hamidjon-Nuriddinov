package redisDB

import (
	"context"
	"homework/model"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func SaveRedis(ctx context.Context, msg model.Company) {
	rdb := Connect()
	if MinValue(ctx, rdb, msg.Company) > msg.Price {
		err := rdb.HSet(ctx, msg.Company, "MinPrice", msg.Price).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if MaxValue(ctx, rdb, msg.Company) < msg.Price {
		err := rdb.HSet(ctx, msg.Company, "MaxPrice", msg.Price).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Getredis(ctx context.Context, companyName string) model.RespCompany {
	rdb := Connect()
	resp := model.RespCompany{
		Company:  companyName,
		MaxPrice: MaxValue(ctx, rdb, companyName),
		MinPrice: MinValue(ctx, rdb, companyName),
	}
	return resp
}

func MaxValue(ctx context.Context, rdb *redis.Client, companyName string) int {
	price := rdb.HGet(ctx, companyName, "MaxPrice").Val()
	if price == "" {
		return 0
	}
	p, err := strconv.Atoi(price)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func MinValue(ctx context.Context, rdb *redis.Client, companyName string) int {
	price := rdb.HGet(ctx, companyName, "MinPrice").Val()
	if price == "" {
		return 10000000
	}
	p, err := strconv.Atoi(price)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
