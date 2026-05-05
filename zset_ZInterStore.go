package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZInterStore(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZInterStore(ctx, "redis_zset_inter", &redis.ZStore{
		Keys: []string{redisZSetKey, redisZSetKey2},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	results, err := rdb.ZRangeWithScores(ctx, "redis_zset_inter", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Member (with score): %v\n", results)
}
