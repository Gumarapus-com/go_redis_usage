package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZRevRange(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZRevRange(ctx, redisZSetKey, 0, 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)

	result2, err := rdb.ZRevRangeWithScores(ctx, redisZSetKey, 0, 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (with scores): %v\n", result2)
}
