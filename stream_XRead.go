package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXRead(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{redisStreamKey},
		Count:   1,
		ID:      "0-0",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
