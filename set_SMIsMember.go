package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSMIsMember(rdb *redis.Client, ctx context.Context) {
	rdb.SAdd(ctx, redisSetKey, "a", "b", "c")

	result, err := rdb.SMIsMember(ctx, redisSetKey, "a", "b", "x").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
