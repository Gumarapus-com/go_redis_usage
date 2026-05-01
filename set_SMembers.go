package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSMembers(rdb *redis.Client, ctx context.Context) {
	// Reset data
	rdb.Del(ctx, redisSetKey)
	rdb.SAdd(ctx, redisSetKey, "a", "b", "c")

	result, err := rdb.SMembers(ctx, redisSetKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
