package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXPending(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.XPending(ctx, redisStreamKey, redisStreamGroup).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %#v\n", result)
}
