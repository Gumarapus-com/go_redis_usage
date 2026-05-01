package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSPop(rdb *redis.Client, ctx context.Context) {
	// Reset data
	rdb.Del(ctx, redisSetKey)
	rdb.SAdd(ctx, redisSetKey, "a", "b", "c", "d", "e")

	result, err := rdb.SPop(ctx, redisSetKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)
}
