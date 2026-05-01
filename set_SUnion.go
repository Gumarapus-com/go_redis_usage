package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSUnion(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisSetKey, redisSetKey2)
	rdb.SAdd(ctx, redisSetKey, "a", "b")
	rdb.SAdd(ctx, redisSetKey2, "b", "c")

	result, err := rdb.SUnion(ctx, redisSetKey, redisSetKey2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
