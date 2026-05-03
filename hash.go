package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	redisHashKey = "redis_hash"
)

func Hash(rdb *redis.Client, ctx context.Context) {
	fmt.Println("HSET")
	runHSet(rdb, ctx)

	fmt.Println("HMSET")
	runHMSet(rdb, ctx)

	fmt.Println("HGET")
	runHGet(rdb, ctx)

	fmt.Println("HGETALL")
	runHGetAll(rdb, ctx)

	fmt.Println("HMGET")
	runHMGet(rdb, ctx)

	fmt.Println("HDEL")
	runHDel(rdb, ctx)

	fmt.Println("HLEN")
	runHLen(rdb, ctx)

	fmt.Println("HKEYS")
	runHKeys(rdb, ctx)

	fmt.Println("HVALS")
	runHVals(rdb, ctx)

	fmt.Println("HEXISTS")
	runHExists(rdb, ctx)

	fmt.Println("HINCRBY")
	runHIncrBy(rdb, ctx)

	fmt.Println("HINCRBYFLOAT")
	runHIncrByFloat(rdb, ctx)
}
