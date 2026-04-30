package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func runBRPop(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey+"_block")

	result, err := rdb.BRPop(ctx, 1*time.Second, redisListKey+"_block").Result()
	if err != nil {
		fmt.Println("  Error: no data retrieved within 1 second ")
	}
	fmt.Printf("  Result: %v\n", result)
}
