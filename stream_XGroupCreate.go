package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXGroupCreate(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.XGroupCreateMkStream(ctx, redisStreamKey, redisStreamGroup+"2", "0").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
