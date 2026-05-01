package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisListKey = "redis_list"

func List(rdb *redis.Client, ctx context.Context) {
	fmt.Println("LPUSH")
	runLPush(rdb, ctx)

	fmt.Println("RPUSH")
	runRPush(rdb, ctx)

	fmt.Println("LPOP")
	runLPop(rdb, ctx)

	fmt.Println("RPOP")
	runRPop(rdb, ctx)

	fmt.Println("LLEN")
	runLLen(rdb, ctx)

	fmt.Println("LRANGE")
	runLRange(rdb, ctx)

	fmt.Println("LINDEX")
	runLIndex(rdb, ctx)

	fmt.Println("LSET")
	runLSet(rdb, ctx)

	fmt.Println("LINSERT")
	runLInsert(rdb, ctx)

	fmt.Println("LREM")
	runLRem(rdb, ctx)

	fmt.Println("LTRIM")
	runLTrim(rdb, ctx)

	fmt.Println("LPUSHX")
	runLPushX(rdb, ctx)

	fmt.Println("RPUSHX")
	runRPushX(rdb, ctx)

	fmt.Println("LMOVE")
	runLMove(rdb, ctx)

	fmt.Println("BLPOP")
	runBLPop(rdb, ctx)

	fmt.Println("BRPOP")
	runBRPop(rdb, ctx)

	fmt.Println("BRPOPLPUSH")
	runBRPopLPush(rdb, ctx)
}
