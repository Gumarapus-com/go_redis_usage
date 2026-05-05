package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func runZUnion(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisZSetKey, redisZSetKey2)
	rdb.ZAdd(ctx, redisZSetKey, redis.Z{Score: 1, Member: "a"}, redis.Z{Score: 2, Member: "b"})
	rdb.ZAdd(ctx, redisZSetKey2, redis.Z{Score: 3, Member: "b"}, redis.Z{Score: 4, Member: "c"})

	result, err := rdb.ZUnion(ctx, redis.ZStore{
		Keys: []string{redisZSetKey, redisZSetKey2},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}
