package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func runBZPopMin(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisZSetKey+"_block")
	rdb.ZAdd(
		ctx,
		redisZSetKey+"_block",
		redis.Z{Score: 1, Member: "a"},
		redis.Z{Score: 2, Member: "b"},
		redis.Z{Score: 3, Member: "c"},
	)

	result, err := rdb.BZPopMin(ctx, time.Second, redisZSetKey+"_block").Result()
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	}
	fmt.Printf("  Result: %v\n", result)
}
