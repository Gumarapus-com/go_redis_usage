package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runSMove(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisSetKey3)

	result, err := rdb.SMove(ctx, redisSetKey, redisSetKey3, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)

	members2, _ := rdb.SMembers(ctx, redisSetKey3).Result()
	fmt.Printf("  Dest members: %v\n", members2)
}
