package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSRandMemberN(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SRandMemberN(ctx, redisSetKey, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
