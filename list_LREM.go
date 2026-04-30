package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLRem(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)
	rdb.RPush(ctx, redisListKey, "a", "b", "a", "c", "a")

	result, err := rdb.LRem(ctx, redisListKey, 2, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	values, _ := rdb.LRange(ctx, redisListKey, 0, -1).Result()
	fmt.Printf("  List values: %v\n", values)
}
