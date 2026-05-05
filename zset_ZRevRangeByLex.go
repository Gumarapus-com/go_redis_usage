package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZRevRangeByLex(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.ZRevRangeByLex(ctx, redisZSetKey, &redis.ZRangeBy{
		Min: "-",
		Max: "+",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (all): %v\n", result)

	result, err = rdb.ZRevRangeByLex(ctx, redisZSetKey, &redis.ZRangeBy{
		Min: "(a",
		Max: "[d",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (a < AND <= d): %v\n", result)

	result, err = rdb.ZRevRangeByLex(ctx, redisZSetKey, &redis.ZRangeBy{
		Min: "[a",
		Max: "(d",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (a <= AND < d): %v\n", result)

	// Alternative
	result, err = rdb.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:   redisZSetKey,
		Start: "[a",
		Stop:  "(d",
		ByLex: true,
		// The reverse option
		Rev: true,
	}).Result()
	fmt.Printf("  Result (using ZRANGE): %v\n", result)
}
