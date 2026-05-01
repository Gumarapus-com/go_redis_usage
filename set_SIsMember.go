package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSIsMember(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SIsMember(ctx, redisSetKey, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (exists): %v\n", result)

	result2, err := rdb.SIsMember(ctx, redisSetKey, "nonexistent").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (not exists): %v\n", result2)
}
