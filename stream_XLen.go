package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXLen(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.XLen(ctx, redisStreamKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
