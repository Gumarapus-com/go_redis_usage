package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZRank(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZRank(ctx, redisZSetKey, "b").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	result2, err := rdb.ZRevRank(ctx, redisZSetKey, "b").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  RevRank: %d\n", result2)
}
