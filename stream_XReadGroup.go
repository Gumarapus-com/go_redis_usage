package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func runXReadGroup(rdb *redis.Client, ctx context.Context) {
	rdb.XGroupCreateMkStream(ctx, redisStreamKey, redisStreamGroup, "0")

	result, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    redisStreamGroup,
		Consumer: redisStreamConsumer,
		Streams:  []string{redisStreamKey, ">"},
		Count:    1,
		Block:    1 * time.Second,
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
