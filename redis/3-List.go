package main

import (
	"context"
	"fmt"
	"github.com/XQ-Gang/learnGo/utils"
	"time"
)

func LearnList() {
	var rdb = utils.RDB
	var ctx = context.Background()

	key := "list-lang"
	value1, value2, value3 := "Python", "Golang", "Java"

	// RPUSH key value1 [value2]
	// 在列表中添加一个或多个值
	// RPUSHX key value
	// 为已存在的列表添加值
	rdb.RPush(ctx, key, value1, value2)

	// LPUSH key value1 [value2]
	// 将一个或多个值插入到列表头部
	// LPUSHX key value
	// 将一个值插入到已存在的列表头部
	rdb.LPush(ctx, key, value3)

	// LLEN key
	// 获取列表长度
	fmt.Println(rdb.LLen(ctx, key)) // 3

	// LRANGE key start stop
	// 获取列表指定范围内的元素
	fmt.Println(rdb.LRange(ctx, key, 1, 2)) // [Python Golang]

	// LINSERT key BEFORE|AFTER pivot value
	// 在列表的元素前或者后插入元素
	rdb.LInsert(ctx, key, "after", "Java", "C")

	// LINDEX key index
	// 通过索引获取列表中的元素
	fmt.Println(rdb.LIndex(ctx, key, 1)) // C

	// LSET key index value
	// 通过索引设置列表元素的值
	rdb.LSet(ctx, key, 1, "C++")

	// LREM key count value
	// 移除列表元素
	// count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT
	// count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值
	// count = 0 : 移除表中所有与 VALUE 相等的值
	rdb.LRem(ctx, key, 1, "Java")
	rdb.LRem(ctx, key, -1, "Java")
	rdb.LRem(ctx, key, 0, "Java")

	// RPOP key
	// 移除列表的最后一个元素，返回值为移除的元素
	// LPOP key
	// 移出并获取列表的第一个元素
	fmt.Println(rdb.RPop(ctx, key)) // Golang
	fmt.Println(rdb.LPop(ctx, key)) // C++

	// RPOPLPUSH source destination
	// 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	rdb.RPopLPush(ctx, key, "list-temp")

	// LTRIM key start stop
	// 对一个列表进行修剪，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
	rdb.LTrim(ctx, "list-temp", 1, -1)

	// BRPOP key1 [key2 ] timeout
	// 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	// BLPOP key1 [key2 ] timeout
	// 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	// BRPOPLPUSH source destination timeout
	// 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它；如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	rdb.BRPop(ctx, time.Second*3, key)

	// 删除测试数据
	// rdb.Del(ctx, key)
}

func main() {
	utils.WrapFunc(LearnList)
}
