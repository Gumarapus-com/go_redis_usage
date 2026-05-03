package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHIncrBy(rdb *redis.Client, ctx context.Context) {
	rdb.HSet(ctx, redisHashKey, "counter", "10")

	result, err := rdb.HIncrBy(ctx, redisHashKey, "counter", 5).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
