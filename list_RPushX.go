package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runRPushX(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)
	rdb.LPush(ctx, redisListKey, "a")

	result, err := rdb.RPushX(ctx, redisListKey, "b").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	values, _ := rdb.LRange(ctx, redisListKey, 0, -1).Result()
	fmt.Printf("  List values: %v\n", values)
}
