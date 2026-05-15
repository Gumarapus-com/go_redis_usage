package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXTrim(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisStreamKey+"_trim")
	for i := 0; i < 10; i++ {
		rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: redisStreamKey + "_trim",
			Values: map[string]any{"num": fmt.Sprint(i)},
		})
	}

	result, err := rdb.XTrimMaxLen(ctx, redisStreamKey+"_trim", 5).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	length, _ := rdb.XLen(ctx, redisStreamKey+"_trim").Result()
	fmt.Printf("  Length after trim: %d\n", length)
}
