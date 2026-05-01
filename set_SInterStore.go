package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSInterStore(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SInterStore(ctx, "redis_set_inter", redisSetKey, redisSetKey2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
