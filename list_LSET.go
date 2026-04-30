package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runLSet(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.LSet(ctx, redisListKey, 0, "updated").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)

	values, _ := rdb.LRange(ctx, redisListKey, 0, -1).Result()
	fmt.Printf("  List values: %v\n", values)
}
