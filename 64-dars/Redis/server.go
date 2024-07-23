package redisDB

import (
	"context"
	"encoding/json"
	"homework/model"
	"log"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

func Server() {
	ctx := context.Background()
	rdb := Connect()
	for {
		Amazon(ctx, rdb, rand.Int()%1000000)
		Google(ctx, rdb, rand.Int()%1000000)
		Meta(ctx, rdb, rand.Int()%1000000)
		Tesla(ctx, rdb, rand.Int()%1000000)
		Netflix(ctx, rdb, rand.Int()%1000000)
		time.Sleep(2 * time.Second)
	}
}

func Amazon(c context.Context, rdb *redis.Client, price int) {
	company := model.Company{
		Company: "Amazon",
		Price:   price,
		Time:    time.Now(),
	}
	data, err := json.Marshal(company)
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Publish(c, company.Company, data).Err()
	if err != nil {
		log.Fatal(err)
	}
	SaveRedis(c, company)

}

func Google(c context.Context, rdb *redis.Client, price int) {
	company := model.Company{
		Company: "Google",
		Price:   price,
		Time:    time.Now(),
	}
	data, err := json.Marshal(company)
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Publish(c, company.Company, data).Err()
	if err != nil {
		log.Fatal(err)
	}
	SaveRedis(c, company)

}

func Meta(c context.Context, rdb *redis.Client, price int) {
	company := model.Company{
		Company: "Meta",
		Price:   price,
		Time:    time.Now(),
	}
	data, err := json.Marshal(company)
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Publish(c, company.Company, data).Err()
	if err != nil {
		log.Fatal(err)
	}
	SaveRedis(c, company)

}

func Netflix(c context.Context, rdb *redis.Client, price int) {
	company := model.Company{
		Company: "Netflix",
		Price:   price,
		Time:    time.Now(),
	}
	data, err := json.Marshal(company)
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Publish(c, company.Company, data).Err()
	if err != nil {
		log.Fatal(err)
	}
	SaveRedis(c, company)
}

func Tesla(c context.Context, rdb *redis.Client, price int) {
	company := model.Company{
		Company: "Tesla",
		Price:   price,
		Time:    time.Now(),
	}
	data, err := json.Marshal(company)
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Publish(c, company.Company, data).Err()
	if err != nil {
		log.Fatal(err)
	}
	SaveRedis(c, company)
}
