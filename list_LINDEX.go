package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLIndex(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)
	rdb.RPush(ctx, redisListKey, "a", "b", "c")

	result, err := rdb.LIndex(ctx, redisListKey, 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (index 0): %s\n", result)

	result2, err := rdb.LIndex(ctx, redisListKey, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (index -1): %s\n", result2)
}
