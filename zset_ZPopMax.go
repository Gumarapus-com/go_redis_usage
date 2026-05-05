package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZPopMax(rdb *redis.Client, ctx context.Context) {
	// Reset data
	rdb.Del(ctx, redisZSetKey)
	rdb.ZAdd(
		ctx,
		redisZSetKey,
		redis.Z{Score: 1, Member: "a"},
		redis.Z{Score: 2, Member: "b"},
		redis.Z{Score: 3, Member: "c"},
	)

	result, err := rdb.ZPopMax(ctx, redisZSetKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
