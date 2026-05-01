package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSCard(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SCard(ctx, redisSetKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
