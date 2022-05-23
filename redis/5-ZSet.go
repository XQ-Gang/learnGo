package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func LearnZSet() {
	key1, key2 := "zset-lang1", "zset-lang2"
	value1, value2, value3, value4 := "Python", "Golang", "Java", "C"

	// ZADD key score1 member1 [score2 member2]
	// 向有序集合添加一个或多个成员，或者更新已存在成员的分数
	rdb.ZAdd(ctx, key1,
		&redis.Z{10, value1},
		&redis.Z{8, value2},
		&redis.Z{6, value3})
	rdb.ZAdd(ctx, key2,
		&redis.Z{9, value2},
		&redis.Z{7, value3},
		&redis.Z{5, value4})

	// ZCARD key
	// 获取有序集合的成员数
	fmt.Println(rdb.ZCard(ctx, key1)) // 3

	// ZINCRBY key increment member
	// 有序集合中对指定成员的分数加上增量 increment
	rdb.ZIncrBy(ctx, key1, 2, value1)

	// ZSCORE key member
	// 返回有序集中，成员的分数值
	fmt.Println(rdb.ZScore(ctx, key1, value1)) // 12

	// ZCOUNT key min max
	// 计算在有序集合中指定区间分数的成员数
	fmt.Println(rdb.ZCount(ctx, key1, "8", "10")) // 1

	// ZRANGE key start stop [WITHSCORES]
	// 通过索引区间返回有序集合指定区间内的成员
	// ZRANGEBYLEX key min max [LIMIT offset count]
	// 通过字典区间返回有序集合的成员
	// ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT]
	// 通过分数返回有序集合指定区间内的成员
	// ZREVRANGE key start stop [WITHSCORES]
	// 返回有序集中指定区间内的成员，通过索引，分数从高到低
	// ZREVRANGEBYSCORE key max min [WITHSCORES]
	// 返回有序集中指定分数区间内的成员，分数从高到低排序
	fmt.Println(rdb.ZRange(ctx, key1, 1, 2))    // [Golang Python]
	fmt.Println(rdb.ZRevRange(ctx, key1, 1, 2)) // [Golang Java]
	fmt.Println(rdb.ZRangeByScore(ctx, key1,
		&redis.ZRangeBy{Min: "6", Max: "12"})) // [Java Golang Python]
	fmt.Println(rdb.ZRevRangeByScore(ctx, key1,
		&redis.ZRangeBy{Min: "6", Max: "12"})) // [Python Golang Java]

	// ZRANK key member
	// 返回有序集合中指定成员的索引
	// ZREVRANK key member
	// 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	fmt.Println(rdb.ZRank(ctx, key1, value3))    // 0
	fmt.Println(rdb.ZRevRank(ctx, key1, value3)) // 2

	// ZREM key member [member ...]
	// 移除有序集合中的一个或多个成员
	// ZREMRANGEBYLEX key min max
	// 移除有序集合中给定的字典区间的所有成员
	// ZREMRANGEBYRANK key start stop
	// 移除有序集合中给定的排名区间的所有成员
	// ZREMRANGEBYSCORE key min max
	// 移除有序集合中给定的分数区间的所有成员
	rdb.ZRem(ctx, key2, value2)                // Remove Golang
	rdb.ZRemRangeByRank(ctx, key2, 1, 1)       // Remove Java
	rdb.ZRemRangeByScore(ctx, key2, "1", "10") // Remove C

	// ZSCAN key cursor [MATCH pattern] [COUNT count]
	// 迭代有序集合中的元素（包括元素成员和元素分值）
	fmt.Println(rdb.ZScan(ctx, key1, 0, "*o*", 2)) // [Golang 8 Python 12]
}
