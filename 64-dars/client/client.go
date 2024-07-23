package main

import (
	"context"
	"fmt"
	redisDB "homework/Redis"
	"log"
)

func main() {
	ctx := context.Background()
	rdb := redisDB.Connect()

	pubsub := rdb.Subscribe(ctx, "Amazon", "Google", "Meta", "Tesla", "Netflix")
	defer pubsub.Close()

	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg, "\n\n")
	}
}
