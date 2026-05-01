package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSDiffStore(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SDiffStore(ctx, "redis_set_diff", redisSetKey, redisSetKey2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
