package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	redisSetKey  = "redis_set"
	redisSetKey2 = "redis_set_2"
	redisSetKey3 = "redis_set_3"
)

func Set(rdb *redis.Client, ctx context.Context) {
	fmt.Println("SADD")
	runSAdd(rdb, ctx)

	fmt.Println("SREM")
	runSRem(rdb, ctx)

	fmt.Println("SMEMBERS")
	runSMembers(rdb, ctx)

	fmt.Println("SCARD")
	runSCard(rdb, ctx)

	fmt.Println("SISMEMBER")
	runSIsMember(rdb, ctx)

	fmt.Println("SMISMEMBER")
	runSMIsMember(rdb, ctx)

	fmt.Println("SUNION")
	runSUnion(rdb, ctx)

	fmt.Println("SUNIONSTORE")
	runSUnionStore(rdb, ctx)

	fmt.Println("SINTER")
	runSInter(rdb, ctx)

	fmt.Println("SINTERSTORE")
	runSInterStore(rdb, ctx)

	fmt.Println("SDIFF")
	runSDiff(rdb, ctx)

	fmt.Println("SDIFFSTORE")
	runSDiffStore(rdb, ctx)

	fmt.Println("SPOP")
	runSPop(rdb, ctx)

	fmt.Println("SPOPN")
	runSPopN(rdb, ctx)

	fmt.Println("SRANDMEMBER")
	runSRandMember(rdb, ctx)

	fmt.Println("SRANDMEMBERN")
	runSRandMemberN(rdb, ctx)

	fmt.Println("SMOVE")
	runSMove(rdb, ctx)
}
