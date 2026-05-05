package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZRange(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZRange(ctx, redisZSetKey, 0, 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)

	result2, err := rdb.ZRangeWithScores(ctx, redisZSetKey, 1, 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (with score): %v\n", result2)
}
