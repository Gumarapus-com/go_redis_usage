package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runXClaim(rdb *redis.Client, ctx context.Context) {
	// Reset
	rdb.Del(ctx, redisStreamKey)
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: redisStreamKey,
		Values: map[string]any{"field": "value"},
	})

	// Create the consumer group first
	rdb.XGroupCreate(ctx, redisStreamKey, redisStreamGroup, "0")
	// `consumer-a` read message (marks them as pending)
	msg_id, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    redisStreamGroup,
		Consumer: "consumer-a",
		Count:    1,
		Streams:  []string{redisStreamKey, ">"},
	}).Result()
	if err != nil {
		panic(err)
	}

	// Claim the message as `consumer-b`
	result, err := rdb.XClaim(ctx, &redis.XClaimArgs{
		Stream:   redisStreamKey,
		Group:    redisStreamGroup,
		Consumer: "consumer-b",
		Messages: []string{msg_id[0].Messages[0].ID},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
