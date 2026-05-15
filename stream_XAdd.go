package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXAdd(rdb *redis.Client, ctx context.Context) {
	// Reset data
	rdb.Del(ctx, redisStreamKey)

	result, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: redisStreamKey,
		Values: map[string]any{"field": "value1"},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)

	entries, _ := rdb.XRange(ctx, redisStreamKey, "-", "+").Result()
	fmt.Printf("  Entries: %v\n", entries)
}
