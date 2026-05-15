package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func runXRange(rdb *redis.Client, ctx context.Context) {
	// Reset data
	rdb.Del(ctx, redisStreamKey)

	// The timestamp (by millisecond) of `now - 10 minutes`
	min10minutes := time.Now().Add(-10 * time.Minute).UnixMilli()
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: redisStreamKey,
		ID:     fmt.Sprintf("%d-0", min10minutes),
		Values: map[string]string{"field3": "value1", "field4": "value2"},
	})
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: redisStreamKey,
		Values: map[string]string{"field1": "value1", "field2": "value2"},
	})

	result, err := rdb.XRange(ctx, redisStreamKey, "-", "+").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (all): %v\n", result)

	currentTime := time.Now().UnixMilli()
	// 5 minutes ago
	min5minutes := time.Now().Add(-5 * time.Minute).UnixMilli()

	result, err = rdb.XRange(
		ctx,
		redisStreamKey,
		fmt.Sprintf("%d-0", min5minutes),
		fmt.Sprintf("%d-0", currentTime),
	).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
