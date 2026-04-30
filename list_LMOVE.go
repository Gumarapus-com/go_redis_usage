package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLMove(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisListKey, redisListKey+"_dest")
	rdb.LPush(ctx, redisListKey, "a", "b", "c")

	result, err := rdb.LMove(ctx, redisListKey, redisListKey+"_dest", "LEFT", "RIGHT").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)

	srcVals, _ := rdb.LRange(ctx, redisListKey, 0, -1).Result()
	dstVals, _ := rdb.LRange(ctx, redisListKey+"_dest", 0, -1).Result()
	fmt.Printf("  Source: %v, Dest: %v\n", srcVals, dstVals)
}
