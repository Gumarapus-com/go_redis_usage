package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHIncrByFloat(rdb *redis.Client, ctx context.Context) {
	rdb.HSet(ctx, redisHashKey, "score", "10.5")

	result, err := rdb.HIncrByFloat(ctx, redisHashKey, "score", 2.3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %f\n", result)
}
