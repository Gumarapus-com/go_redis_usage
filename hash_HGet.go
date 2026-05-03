package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHGet(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.HGet(ctx, redisHashKey, "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)
}
