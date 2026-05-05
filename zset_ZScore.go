package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZScore(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZScore(ctx, redisZSetKey, "b").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %f\n", result)
}
