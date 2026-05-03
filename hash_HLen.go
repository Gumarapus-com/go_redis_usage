package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHLen(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.HLen(ctx, redisHashKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
