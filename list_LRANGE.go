package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLRange(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)
	rdb.LPush(ctx, redisListKey, "a", "b", "c", "d", "e")

	result, err := rdb.LRange(ctx, redisListKey, 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)

	result2, err := rdb.LRange(ctx, redisListKey, -2, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Last 2: %v\n", result2)
}
