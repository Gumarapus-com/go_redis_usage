package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZRem(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZRem(ctx, redisZSetKey, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	members, _ := rdb.ZRangeWithScores(ctx, redisZSetKey, 0, -1).Result()
	fmt.Printf("  Members: %v\n", members)
}
