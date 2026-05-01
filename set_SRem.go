package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSRem(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SRem(ctx, redisSetKey, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	members, _ := rdb.SMembers(ctx, redisSetKey).Result()
	fmt.Printf("  Members: %v\n", members)
}
