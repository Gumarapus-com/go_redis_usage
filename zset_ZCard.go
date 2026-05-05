package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZCard(rdb *redis.Client, ctx context.Context) {
	// Reset data
	rdb.Del(ctx, redisZSetKey)
	rdb.ZAdd(
		ctx,
		redisZSetKey,
		redis.Z{Score: 1, Member: "a"},
		redis.Z{Score: 2, Member: "b"},
		redis.Z{Score: 3, Member: "c"},
		redis.Z{Score: 4, Member: "d"},
		redis.Z{Score: 5, Member: "e"},
	)

	result, err := rdb.ZCard(ctx, redisZSetKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
