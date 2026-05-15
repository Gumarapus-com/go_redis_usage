package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func runXRevRange(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.XRevRange(ctx, redisStreamKey, "+", "-").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (all): %v\n", result)

	currentTime := time.Now().UnixMilli()
	// 5 minutes ago
	min5minutes := time.Now().Add(-5 * time.Minute).UnixMilli()
	result, err = rdb.XRevRange(
		ctx,
		redisStreamKey,
		fmt.Sprintf("%d-0", currentTime),
		fmt.Sprintf("%d-0", min5minutes),
	).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
