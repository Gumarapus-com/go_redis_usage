package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXAddWithID(rdb *redis.Client, ctx context.Context) {
	redisKey := "redis_stream_tmp"

	// Reset data
	rdb.Del(ctx, redisKey)

	result, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: redisKey,
		ID:     "1000000000000-0",
		Values: map[string]any{
			"field": "value2",
		},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)
}
