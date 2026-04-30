package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLLen(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)
	rdb.RPush(ctx, redisListKey, "a", "b", "c")

	result, err := rdb.LLen(ctx, redisListKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
