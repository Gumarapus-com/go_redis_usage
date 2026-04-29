package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisStringKey = "redis_str"
	redisStrNumKey = "redis_str_num"
	redisStrKey1   = "redis_str_1"
	redisStrKey2   = "redis_str_2"
	redisStrKey3   = "redis_str_3"
)

func runSet(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.Set(ctx, redisStringKey, "a string", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)
}

func runGet(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.Get(ctx, redisStringKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)
}

func runAppend(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.Append(ctx, redisStringKey, " + append string").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	newValue, _ := rdb.Get(ctx, redisStringKey).Result()
	fmt.Printf("  New value: %s\n", newValue)
}

func runStrlen(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.StrLen(ctx, redisStringKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}

func runGetRange(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.GetRange(ctx, redisStringKey, 0, 10).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)
}

func runSetRange(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.SetRange(ctx, redisStringKey, 0, "updated + ").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	newValue, _ := rdb.Get(ctx, redisStringKey).Result()
	fmt.Printf("  New value: %s\n", newValue)
}

func runSetNX(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisStringKey)

	result, err := rdb.SetNX(ctx, redisStringKey, "set nx value", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)

	value, _ := rdb.Get(ctx, redisStringKey).Result()
	fmt.Printf("  Value after SetNX: %s\n", value)

	result2, err := rdb.SetNX(ctx, redisStringKey, "already exists", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (should be false): %v\n", result2)
}

func runSetXX(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisStringKey)

	result, err := rdb.SetXX(ctx, redisStringKey, "should fail", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (key not exists, should be false): %v\n", result)

	rdb.Set(ctx, redisStringKey, "exists", 0)

	result2, err := rdb.SetXX(ctx, redisStringKey, "updated with xx", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (key exists, should be true): %v\n", result2)

	value, _ := rdb.Get(ctx, redisStringKey).Result()
	fmt.Printf("  Value after SetXX: %s\n", value)
}

func runMGet(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrKey1, "value1", 0)
	rdb.Set(ctx, redisStrKey2, "value2", 0)
	rdb.Set(ctx, redisStrKey3, "value3", 0)

	result, err := rdb.MGet(ctx, redisStrKey1, redisStrKey2, redisStrKey3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)
}

func runMSet(rdb *redis.Client, ctx context.Context) {
	result, err := rdb.MSet(ctx, redisStrKey1, "mget1", redisStrKey2, "mget2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)

	value1, _ := rdb.Get(ctx, redisStrKey1).Result()
	value2, _ := rdb.Get(ctx, redisStrKey2).Result()
	fmt.Printf("  Values: %s, %s\n", value1, value2)
}

func runMSetNX(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisStrKey1)
	rdb.Del(ctx, redisStrKey2)
	rdb.Del(ctx, redisStrKey3)

	result, err := rdb.MSetNX(ctx, redisStrKey1, "nx1", redisStrKey2, "nx2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %v\n", result)

	result2, err := rdb.MSetNX(ctx, redisStrKey1, "conflict", redisStrKey3, "nx3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (partial fail): %v\n", result2)

	val1, _ := rdb.Get(ctx, redisStrKey1).Result()
	val3, _ := rdb.Get(ctx, redisStrKey3).Result()
	fmt.Printf("  Values: %s, %s\n", val1, val3)
}

func runIncr(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrNumKey, "10", 0)

	result, err := rdb.Incr(ctx, redisStrNumKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	value, _ := rdb.Get(ctx, redisStrNumKey).Result()
	fmt.Printf("  New value: %s\n", value)
}

func runIncrBy(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrNumKey, "10", 0)

	result, err := rdb.IncrBy(ctx, redisStrNumKey, 5).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	value, _ := rdb.Get(ctx, redisStrNumKey).Result()
	fmt.Printf("  New value: %s\n", value)
}

func runIncrByFloat(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrNumKey, "10.5", 0)

	result, err := rdb.IncrByFloat(ctx, redisStrNumKey, 2.3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)

	value, _ := rdb.Get(ctx, redisStrNumKey).Result()
	fmt.Printf("  New value: %s\n", value)
}

func runDecr(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrNumKey, "10", 0)

	result, err := rdb.Decr(ctx, redisStrNumKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	value, _ := rdb.Get(ctx, redisStrNumKey).Result()
	fmt.Printf("  New value: %s\n", value)
}

func runDecrBy(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrNumKey, "10", 0)

	result, err := rdb.DecrBy(ctx, redisStrNumKey, 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)

	value, _ := rdb.Get(ctx, redisStrNumKey).Result()
	fmt.Printf("  New value: %s\n", value)
}

func runGetSet(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStringKey, "old value", 0)

	result, err := rdb.GetSet(ctx, redisStringKey, "new value").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Old value: %s\n", result)

	newValue, _ := rdb.Get(ctx, redisStringKey).Result()
	fmt.Printf("  New value: %s\n", newValue)
}

func runGetEx(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStringKey, "expires soon", 5*time.Second)

	result, err := rdb.GetEx(ctx, redisStringKey, 10*time.Second).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (with new expiration): %s\n", result)
}

func runGetDel(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStringKey, "to be deleted", 0)

	result, err := rdb.GetDel(ctx, redisStringKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (deleted value): %s\n", result)

	exists, _ := rdb.Exists(ctx, redisStringKey).Result()
	fmt.Printf("  Exists after GetDel: %d\n", exists)
}

func runGetBit(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStringKey, "a", 0)

	result, err := rdb.GetBit(ctx, redisStringKey, 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}

func runSetBit(rdb *redis.Client, ctx context.Context) {
	rdb.Del(ctx, redisStringKey)

	result, err := rdb.SetBit(ctx, redisStringKey, 0, 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (old value): %d\n", result)
}

func runBitCount(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStringKey, "a", 0)

	result, err := rdb.BitCount(ctx, redisStringKey, nil).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}

func runBitOpAnd(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrKey1, "a", 0)
	rdb.Set(ctx, redisStrKey2, "b", 0)

	result, err := rdb.BitOpAnd(ctx, "bit_result", redisStrKey1, redisStrKey2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (dest key): %d\n", result)
}

func runBitOpOr(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrKey1, "a", 0)
	rdb.Set(ctx, redisStrKey2, "b", 0)

	result, err := rdb.BitOpOr(ctx, "bit_result", redisStrKey1, redisStrKey2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}

func runBitOpXor(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStrKey1, "a", 0)
	rdb.Set(ctx, redisStrKey2, "a", 0)

	result, err := rdb.BitOpXor(ctx, "bit_result", redisStrKey1, redisStrKey2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result (same values = 0): %d\n", result)
}

func runBitOpNot(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStringKey, "a", 0)

	result, err := rdb.BitOpNot(ctx, "bit_not_result", redisStringKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}

func runBitPos(rdb *redis.Client, ctx context.Context) {
	rdb.Set(ctx, redisStringKey, "a", 0)

	result, err := rdb.BitPos(ctx, redisStringKey, 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %d\n", result)
}

func runSetWithExp(rdb *redis.Client, ctx context.Context) {
	// This data will be expired in One hour
	result, err := rdb.Set(ctx, redisStringKey, "a string", time.Hour).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  Result: %s\n", result)
}

func String(rdb *redis.Client, ctx context.Context) {
	fmt.Println("SET")
	runSet(rdb, ctx)
	fmt.Println("SET (with EX)")
	runSetWithExp(rdb, ctx)

	fmt.Println("GET")
	runGet(rdb, ctx)

	fmt.Println("APPEND")
	runAppend(rdb, ctx)

	fmt.Println("STRLEN")
	runStrlen(rdb, ctx)

	fmt.Println("GETRANGE")
	runGetRange(rdb, ctx)

	fmt.Println("SETRANGE")
	runSetRange(rdb, ctx)

	fmt.Println("SETNX")
	runSetNX(rdb, ctx)

	fmt.Println("SETXX")
	runSetXX(rdb, ctx)

	fmt.Println("MGET")
	runMGet(rdb, ctx)

	fmt.Println("MSET")
	runMSet(rdb, ctx)

	fmt.Println("MSETNX")
	runMSetNX(rdb, ctx)

	fmt.Println("INCR")
	runIncr(rdb, ctx)

	fmt.Println("INCRBY")
	runIncrBy(rdb, ctx)

	fmt.Println("INCRBYFLOAT")
	runIncrByFloat(rdb, ctx)

	fmt.Println("DECR")
	runDecr(rdb, ctx)

	fmt.Println("DECRBY")
	runDecrBy(rdb, ctx)

	fmt.Println("GETSET")
	runGetSet(rdb, ctx)

	fmt.Println("GETEX")
	runGetEx(rdb, ctx)

	fmt.Println("GETDEL")
	runGetDel(rdb, ctx)

	fmt.Println("GETBIT")
	runGetBit(rdb, ctx)

	fmt.Println("SETBIT")
	runSetBit(rdb, ctx)

	fmt.Println("BITCOUNT")
	runBitCount(rdb, ctx)

	fmt.Println("BITOPAND")
	runBitOpAnd(rdb, ctx)

	fmt.Println("BITOPOR")
	runBitOpOr(rdb, ctx)

	fmt.Println("BITOPXOR")
	runBitOpXor(rdb, ctx)

	fmt.Println("BITOPNOT")
	runBitOpNot(rdb, ctx)

	fmt.Println("BITPOS")
	runBitPos(rdb, ctx)
}
