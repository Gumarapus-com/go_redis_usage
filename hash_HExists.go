package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHExists(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.HExists(ctx, redisHashKey, "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (exists): %v\n", result)

	result2, err := rdb.HExists(ctx, redisHashKey, "nonexistent").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (not exists): %v\n", result2)
}
