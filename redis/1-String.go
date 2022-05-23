package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func LearnString() {
	key, value := "str-name", "XiuQiuGang"
	key2, value2 := "str-age", "20"

	// SET key value
	// 设置指定 key 的值
	// expiration=0 表示没有过期时间
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	err = rdb.Set(ctx, key2, value2, 0).Err()
	if err != nil {
		panic(err)
	}

	// GET key
	// 获取指定 key 的值
	// redis.Nil 表示 key 不存在
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("get str-name:", val) // XiuQiuGang
	}

	// GETRANGE key start end
	// 返回 key 中字符串值的子字符
	// end = -1 表示直到末尾
	fmt.Println(rdb.GetRange(ctx, key, 6, -1)) // Gang

	// GETSET key value
	// 将给定 key 的值设为 value ，并返回 key 的旧值
	fmt.Println(rdb.GetSet(ctx, key, "XQGang")) // XiuQiuGang
	fmt.Println(rdb.GetSet(ctx, key, value))    // XQGang

	// SETEX key seconds value
	// 将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)
	err = rdb.SetEX(ctx, key, value, time.Second*10).Err()
	fmt.Println(rdb.TTL(ctx, key)) // 10s

	// SETNX key value
	// 只有在 key 不存在时设置 key 的值
	ok, err := rdb.SetNX(ctx, key, value, time.Second*10).Result()
	if ok {
		fmt.Println("设置成功")
	} else {
		fmt.Println("key已经存在缓存中，设置失败") // Output
	}

	// STRLEN key
	// 返回 key 所储存的字符串值的长度
	fmt.Println(rdb.StrLen(ctx, key)) // 10

	// MGET key1 [key2..]
	// 获取所有(一个或多个)给定 key 的值。
	fmt.Println(rdb.MGet(ctx, key, key2)) // [XiuQiuGang 20]

	// INCR key
	// 将 key 中储存的数字值增一
	fmt.Println(rdb.Incr(ctx, key2)) // 21

	// INCRBY key increment
	// 将 key 所储存的值加上给定的增量值（increment）
	fmt.Println(rdb.IncrBy(ctx, key2, 5)) // 26

	// DECR key
	// 将 key 中储存的数字值减一
	fmt.Println(rdb.Decr(ctx, key2)) // 25

	// DECRBY key decrement
	// key 所储存的值减去给定的减量值（decrement）
	fmt.Println(rdb.DecrBy(ctx, key2, 5)) // 20
}
