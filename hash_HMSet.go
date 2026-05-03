package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHMSet(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.HMSet(ctx, redisHashKey, map[string]interface{}{
		"field4": "value4",
		"field5": "value5",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
