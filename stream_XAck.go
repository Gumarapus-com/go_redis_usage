package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXAck(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisStreamKey+"_ack")
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: redisStreamKey + "_ack",
		Values: map[string]any{"field": "value"},
	})

	rdb.XGroupCreateMkStream(ctx, redisStreamKey+"_ack", redisStreamGroup, "0").Result()

	result, err := rdb.XAck(ctx, redisStreamKey+"_ack", redisStreamGroup, "0-0").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}
