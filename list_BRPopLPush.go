package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func runBRPopLPush(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey, redisListKey+"_dest")
	rdb.LPush(ctx, redisListKey, "a")

	result, err := rdb.BRPopLPush(ctx, redisListKey, redisListKey+"_dest", 1*time.Second).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)

	dstVals, _ := rdb.LRange(ctx, redisListKey+"_dest", 0, -1).Result()
	fmt.Printf("  Dest values: %v\n", dstVals)
}
