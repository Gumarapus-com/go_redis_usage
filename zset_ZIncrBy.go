package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZIncrBy(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZIncrBy(ctx, redisZSetKey, 2, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %f\n", result)

	score, _ := rdb.ZScore(ctx, redisZSetKey, "a").Result()
	fmt.Printf("  New score: %f\n", score)
}
