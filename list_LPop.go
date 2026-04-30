package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLPop(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)
	rdb.RPush(ctx, redisListKey, "a", "b", "c")

	result, err := rdb.LPop(ctx, redisListKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)

	values, _ := rdb.LRange(ctx, redisListKey, 0, -1).Result()
	fmt.Printf("  List values: %v\n", values)
}
