package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZInter(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZInter(ctx, &redis.ZStore{
		Keys: []string{redisZSetKey, redisZSetKey2},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
