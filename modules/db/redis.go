package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func initRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		panic(">>> connect redis failed, error=" + err.Error())
	}
	fmt.Println(">>> connect redis success")
}
