package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZRangeByScore(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZRangeByScore(ctx, redisZSetKey, &redis.ZRangeBy{
		Min: "1",
		Max: "2",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)

	// Alternative: using ZRANGE
	result, err = rdb.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:     redisZSetKey,
		Start:   1,
		Stop:    2,
		ByScore: true,
	}).Result()
	fmt.Printf("  Result (using ZRANGE): %v\n", result)
}
