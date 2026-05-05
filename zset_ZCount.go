package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZCount(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZCount(ctx, redisZSetKey, "1", "3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
