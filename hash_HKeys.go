package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHKeys(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.HKeys(ctx, redisHashKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
