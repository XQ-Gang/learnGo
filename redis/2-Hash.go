package main

import (
	"context"
	"fmt"
	"github.com/XQ-Gang/learnGo/utils"
	"github.com/go-redis/redis/v8"
)

func LearnHash() {
	var rdb = utils.RDB
	var ctx = context.Background()

	key := "hash-gang"
	key1, value1 := "name", "XiuQiuGang"
	key2, value2 := "age", "20"
	values := map[string]string{key1: value1, key2: value2}

	// HSET key field value
	// 将哈希表 key 中的字段 field 的值设为 value
	rdb.HSet(ctx, key, key1, value1)
	rdb.HSet(ctx, key, key2, value2)

	// HDEL key field1 [field2]
	// 删除一个或多个哈希表字段
	rdb.HDel(ctx, key, key2)

	// HEXISTS key field
	// 查看哈希表 key 中，指定的字段是否存在
	fmt.Println(rdb.HExists(ctx, key, key1)) // true
	fmt.Println(rdb.HExists(ctx, key, key2)) // false

	// HGET key field
	// 获取存储在哈希表中指定字段的值
	fmt.Println(rdb.HGet(ctx, key, key1)) // XiuQiuGang
	age, err := rdb.HGet(ctx, key, key2).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist") // Output
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("hget hash-gang age:", age)
	}

	// HMSET key field1 value1 [field2 value2 ]
	// 同时将多个 field-value (域-值)对设置到哈希表 key 中
	rdb.HMSet(ctx, key, values)

	// HMGET key field1 [field2]
	// 获取所有给定字段的值
	fmt.Println(rdb.HMGet(ctx, key, key1, key2)) // [XiuQiuGang 20]

	// HGETALL key
	// 获取在哈希表中指定 key 的所有字段和值
	fmt.Println(rdb.HGetAll(ctx, key)) // map[age:20 gender:male name:XiuQiuGang]

	// HLEN key
	// 获取哈希表中字段的数量
	fmt.Println(rdb.HLen(ctx, key)) // 3

	// HKEYS key
	// 获取所有哈希表中的字段
	fmt.Println(rdb.HKeys(ctx, key)) // [name gender age]

	// HVALS key
	// 获取哈希表中所有值
	fmt.Println(rdb.HVals(ctx, key)) // [XiuQiuGang male 20]

	// HINCRBY key field increment
	// 为哈希表 key 中的指定字段的整数值加上增量 increment
	fmt.Println(rdb.HIncrBy(ctx, key, key2, 10)) // 30
	fmt.Println(rdb.HIncrBy(ctx, key, key2, -5)) // 25

	// HSETNX key field value
	// 只有在字段 field 不存在时，设置哈希表字段的值
	fmt.Println(rdb.HSetNX(ctx, key, key1, "XQGang"))   // false
	fmt.Println(rdb.HSetNX(ctx, key, "gender", "male")) // true

	// HSCAN key cursor [MATCH pattern] [COUNT count]
	// 迭代哈希表中的键值对
	fmt.Println(rdb.HScan(ctx, key, 0, "name", 1)) // [name XiuQiuGang]

	// 删除测试数据
	// rdb.Del(ctx, key)
}

func main() {
	utils.WrapFunc(LearnHash)
}
