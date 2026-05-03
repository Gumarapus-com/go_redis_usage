package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHMGet(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.HMGet(ctx, redisHashKey, "field1", "field2", "notexist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
