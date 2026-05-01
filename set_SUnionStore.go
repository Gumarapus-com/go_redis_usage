package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSUnionStore(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SUnionStore(ctx, "redis_set_union", redisSetKey, redisSetKey2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	newSetMembers, _ := rdb.SMembers(ctx, "redis_set_union").Result()
	fmt.Printf("  Result: %v\n", newSetMembers)
}
