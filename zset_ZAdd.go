package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZAdd(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisZSetKey)

	result, err := rdb.ZAdd(
		ctx,
		redisZSetKey,
		redis.Z{Score: 1, Member: "a"},
		redis.Z{Score: 2, Member: "b"},
		redis.Z{Score: 3, Member: "c"},
	).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	members, _ := rdb.ZRangeWithScores(ctx, redisZSetKey, 0, -1).Result()
	fmt.Printf("  Members: %v\n", members)
}
