package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLInsert(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey)
	rdb.RPush(ctx, redisListKey, "a", "c")

	result, err := rdb.LInsertBefore(ctx, redisListKey, "c", "b").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (before): %d\n", result)

	result, err = rdb.LInsertAfter(ctx, redisListKey, "c", "b").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result  (after): %d\n", result)

	values, _ := rdb.LRange(ctx, redisListKey, 0, -1).Result()
	fmt.Printf("  List values: %v\n", values)
}
