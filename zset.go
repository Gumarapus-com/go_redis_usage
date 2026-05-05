package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	redisZSetKey  = "redis_zset"
	redisZSetKey2 = "redis_zset_2"
)

func SortedSet(rdb *redis.Client, ctx context.Context) {
	fmt.Println("ZADD")
	runZAdd(rdb, ctx)

	fmt.Println("ZREM")
	runZRem(rdb, ctx)

	fmt.Println("ZCARD")
	runZCard(rdb, ctx)

	fmt.Println("ZSCORE")
	runZScore(rdb, ctx)

	fmt.Println("ZRANK")
	runZRank(rdb, ctx)

	fmt.Println("ZRANGE")
	runZRange(rdb, ctx)

	fmt.Println("ZRANGEBYSCORE")
	runZRangeByScore(rdb, ctx)

	fmt.Println("ZRANGEBYLEX")
	runZRangeByLex(rdb, ctx)

	fmt.Println("ZREVRANGE")
	runZRevRange(rdb, ctx)

	fmt.Println("ZREVRANGEBYSCORE")
	runZRevRangeByScore(rdb, ctx)

	fmt.Println("ZREVRANGEBYLEX")
	runZRevRangeByLex(rdb, ctx)

	fmt.Println("ZCOUNT")
	runZCount(rdb, ctx)

	fmt.Println("ZINCRBY")
	runZIncrBy(rdb, ctx)

	fmt.Println("ZUNION")
	runZUnion(rdb, ctx)

	fmt.Println("ZUNIONSTORE")
	runZUnionStore(rdb, ctx)

	fmt.Println("ZINTER")
	runZInter(rdb, ctx)

	fmt.Println("ZINTERSTORE")
	runZInterStore(rdb, ctx)

	fmt.Println("ZPOPMIN")
	runZPopMin(rdb, ctx)

	fmt.Println("ZPOPMAX")
	runZPopMax(rdb, ctx)

	fmt.Println("BZPOPMIN")
	runBZPopMin(rdb, ctx)
}
