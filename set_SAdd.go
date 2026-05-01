package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSAdd(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisSetKey)

	result, err := rdb.SAdd(ctx, redisSetKey, "a", "b", "c").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	members, _ := rdb.SMembers(ctx, redisSetKey).Result()
	fmt.Printf("  Members: %v\n", members)
}
