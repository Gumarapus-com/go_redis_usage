package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSPopN(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SPopN(ctx, redisSetKey, 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
