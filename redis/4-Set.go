package main

import (
	"context"
	"fmt"
	"github.com/XQ-Gang/learnGo/utils"
)

func LearnSet() {
	var rdb = utils.RDB
	var ctx = context.Background()

	key1, key2, key3 := "set-lang1", "set-lang2", "set-lang3"
	value1, value2, value3, value4 := "Python", "Golang", "Java", "C"

	// SADD key member1 [member2]
	// 向集合添加一个或多个成员
	rdb.SAdd(ctx, key1, value1, value2, value3)
	rdb.SAdd(ctx, key2, value2, value3, value4)

	// SCARD key
	// 获取集合的成员数
	fmt.Println(rdb.SCard(ctx, key1)) // 3

	// SMEMBERS key
	// 返回集合中的所有成员
	fmt.Println(rdb.SMembers(ctx, key1)) // [Golang Python Java]

	// SISMEMBER key member
	// 判断 member 元素是否是集合 key 的成员
	fmt.Println(rdb.SIsMember(ctx, key1, "C")) // false

	// SDIFF key1 [key2]
	// 返回第一个集合与其他集合之间的差异
	// SDIFFSTORE destination key1 [key2]
	// 返回给定所有集合的差集并存储在 destination 中
	fmt.Println(rdb.SDiff(ctx, key1, key2)) // [Python]

	// SINTER key1 [key2]
	// 返回给定所有集合的交集
	// SINTERSTORE destination key1 [key2]
	// 返回给定所有集合的交集并存储在 destination 中
	fmt.Println(rdb.SInter(ctx, key1, key2)) // [Golang Java]

	// SUNION key1 [key2]
	// 返回所有给定集合的并集
	fmt.Println(rdb.SUnion(ctx, key1, key2)) // [Python Golang C Java]

	// SMOVE source destination member
	// 将 member 元素从 source 集合移动到 destination 集合
	rdb.SMove(ctx, key2, key3, "C")
	rdb.SMove(ctx, key2, key3, "R") // 无影响

	// SRANDMEMBER key [count]
	// 返回集合中一个或多个随机数
	fmt.Println(rdb.SRandMember(ctx, key1)) // Python

	// SPOP key
	// 移除并返回集合中的一个随机元素
	fmt.Println(rdb.SPop(ctx, key3)) // C

	// SREM key member1 [member2]
	// 移除集合中一个或多个成员
	rdb.SRem(ctx, key2, "Golang", "Java")

	// SSCAN key cursor [MATCH pattern] [COUNT count]
	// 迭代集合中的元素
	fmt.Println(rdb.SScan(ctx, key1, 0, "*", 2).Result()) // [Golang Python] 2 <nil>
	fmt.Println(rdb.SScan(ctx, key1, 0, "*", 4).Result()) // [Golang Python Java] 0 <nil>

	// 删除测试数据
	// rdb.Del(ctx, key1, key2, key3)
}

func main() {
	utils.WrapFunc(LearnSet)
}
