package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXGroupSetID(rdb *redis.Client, ctx context.Context) {
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: redisStreamKey,
		Values: map[string]any{"field": "value4"},
	})

	result, err := rdb.XGroupSetID(ctx, redisStreamKey, redisStreamGroup, "0").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
