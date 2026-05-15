package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	redisStreamKey      = "redis_stream"
	redisStreamGroup    = "redis_stream_group"
	redisStreamConsumer = "redis_stream_consumer"
)

func Stream(rdb *redis.Client, ctx context.Context) {
	fmt.Println("XADD")
	runXAdd(rdb, ctx)

	fmt.Println("XADD (with ID)")
	runXAddWithID(rdb, ctx)

	fmt.Println("XLEN")
	runXLen(rdb, ctx)

	fmt.Println("XRANGE")
	runXRange(rdb, ctx)

	fmt.Println("XREVRANGE")
	runXRevRange(rdb, ctx)

	fmt.Println("XREAD")
	runXRead(rdb, ctx)

	fmt.Println("XREADGROUP")
	runXReadGroup(rdb, ctx)

	fmt.Println("XGROUPCREATE")
	runXGroupCreate(rdb, ctx)

	fmt.Println("XGROUPSETID")
	runXGroupSetID(rdb, ctx)

	fmt.Println("XACK")
	runXAck(rdb, ctx)

	fmt.Println("XPENDING")
	runXPending(rdb, ctx)

	fmt.Println("XCLAIM")
	runXClaim(rdb, ctx)

	fmt.Println("XTRIM")
	runXTrim(rdb, ctx)
}
