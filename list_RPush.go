package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runRPush(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)

	result, err := rdb.RPush(ctx, redisListKey, "a", "b", "c").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	values, _ := rdb.LRange(ctx, redisListKey, 0, -1).Result()
	fmt.Printf("  List values: %v\n", values)
}
