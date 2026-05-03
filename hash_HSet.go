package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runHSet(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisHashKey)

	result, err := rdb.HSet(ctx, redisHashKey, map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	values, _ := rdb.HGetAll(ctx, redisHashKey).Result()
	fmt.Printf("  Values: %v\n", values)
}
